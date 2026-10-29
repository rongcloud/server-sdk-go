package sdk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestMessageMarshalDefaults(t *testing.T) {
	for _, tt := range []struct {
		name      string
		options   string
		persisted int
		counted   int
	}{
		{name: "defaults", options: `{}`, persisted: 1, counted: 1},
		{name: "disabled", options: `{"isPersisted":0,"isCounted":0}`},
		{name: "enabled", options: `{"isPersisted":1,"isCounted":1}`, persisted: 1, counted: 1},
		{name: "default counted", options: `{"isPersisted":0}`, counted: 1},
		{name: "default persisted", options: `{"isCounted":0}`, persisted: 1},
		{name: "mixed", options: `{"isPersisted":1,"isCounted":0}`, persisted: 1},
		{name: "null defaults", options: `{"isPersisted":null,"isCounted":null}`, persisted: 1, counted: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			msg := Message{Content: `{"content":"hello"}`, ObjectName: "RC:TxtMsg", DisableUpdateLastMsg: true}
			if err := json.Unmarshal([]byte(tt.options), &msg); err != nil {
				t.Fatal(err)
			}
			original := msg
			body, err := json.Marshal(msg)
			if err != nil {
				t.Fatal(err)
			}
			var payload map[string]interface{}
			if err := json.Unmarshal(body, &payload); err != nil {
				t.Fatal(err)
			}
			if payload["isPersisted"] != float64(tt.persisted) || payload["isCounted"] != float64(tt.counted) {
				t.Errorf("message = %s, want isPersisted=%d, isCounted=%d", body, tt.persisted, tt.counted)
			}
			if payload["content"] != msg.Content || payload["objectName"] != msg.ObjectName || payload["disableUpdateLastMsg"] != true {
				t.Errorf("existing message fields changed: %s", body)
			}
			if !reflect.DeepEqual(msg, original) {
				t.Error("marshaling changed the original message")
			}
		})
	}
}

type pushMessageTransport func(*http.Request) (*http.Response, error)

func (f pushMessageTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestPushSendMessageOptions(t *testing.T) {
	for _, tt := range []struct {
		name      string
		options   string
		persisted int
		counted   int
		pushOnly  bool
	}{
		{name: "broadcast defaults", options: `{}`, persisted: 1, counted: 1},
		{name: "broadcast disabled", options: `{"isPersisted":0,"isCounted":0}`},
		{name: "broadcast mixed", options: `{"isPersisted":0}`, counted: 1},
		{name: "push", pushOnly: true},
	} {
		for _, pointer := range []bool{false, true} {
			name := tt.name + "/value"
			if pointer {
				name = tt.name + "/pointer"
			}
			t.Run(name, func(t *testing.T) {
				broadcast := Broadcast{FromUserID: "sender", Message: Message{Content: "hello", ObjectName: "RC:TxtMsg"}}
				if !tt.pushOnly {
					if err := json.Unmarshal([]byte(tt.options), &broadcast.Message); err != nil {
						t.Fatal(err)
					}
				}
				original := broadcast.Message
				var sender Sender = broadcast
				if pointer {
					sender = &broadcast
				}
				if tt.pushOnly {
					push := Push{Notification: Notification{Alert: "hello"}}
					sender = push
					if pointer {
						sender = &push
					}
				}
				calls := 0
				extra := defaultExtra
				rc := &RongCloud{appKey: "test-key", appSecret: "test-secret", rongCloudURI: "https://example.test", rongCloudExtra: &extra}
				rc.SetHttpTransport(pushMessageTransport(func(req *http.Request) (*http.Response, error) {
					calls++
					if req.Method != http.MethodPost || req.URL.Path != "/push.json" {
						t.Errorf("unexpected request: %s %s", req.Method, req.URL)
					}
					var payload map[string]interface{}
					if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
						t.Fatal(err)
					}
					for _, key := range []string{"isPersisted", "isCounted"} {
						if _, ok := payload[key]; ok {
							t.Errorf("%s must be nested inside message", key)
						}
					}
					if tt.pushOnly {
						if _, ok := payload["message"]; ok {
							t.Error("push-only request contains message")
						}
					} else {
						msg, ok := payload["message"].(map[string]interface{})
						if !ok {
							t.Fatal("missing broadcast message")
						}
						if msg["isPersisted"] != float64(tt.persisted) || msg["isCounted"] != float64(tt.counted) {
							t.Errorf("message = %v, want isPersisted=%d, isCounted=%d", msg, tt.persisted, tt.counted)
						}
					}
					return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: ioutil.NopCloser(strings.NewReader(`{"code":200,"id":"push-id"}`))}, nil
				}))
				result, err := rc.PushSend(sender)
				if err != nil {
					t.Fatal(err)
				}
				if calls != 1 || result.Code != 200 || result.ID != "push-id" {
					t.Errorf("calls=%d, result=%+v", calls, result)
				}
				if !reflect.DeepEqual(broadcast.Message, original) {
					t.Error("PushSend changed the original message")
				}
			})
		}
	}
}
