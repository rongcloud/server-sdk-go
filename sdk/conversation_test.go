package sdk

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

type conversationQueryTransport func(*http.Request) (*http.Response, error)

func (f conversationQueryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newConversationQueryClient(transport conversationQueryTransport) *RongCloud {
	extra := defaultExtra
	return &RongCloud{
		appKey: "test-key", appSecret: "test-secret",
		rongCloudURI: "https://example.test", rongCloudExtra: &extra,
		globalTransport: transport,
	}
}

func conversationQueryResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       ioutil.NopCloser(strings.NewReader(body)),
	}
}

func TestRongCloud_ConversationQueryValidation(t *testing.T) {
	client := newConversationQueryClient(func(*http.Request) (*http.Response, error) {
		t.Fatal("invalid parameters must not send a request")
		return nil, nil
	})
	for _, userID := range []string{"", " ", "\t\n"} {
		_, err := client.ConversationQuery(QueryConversationModel{UserId: userID})
		if code, ok := err.(CodeResult); !ok || code.Code != 1002 {
			t.Errorf("userId %q: expected parameter error 1002, got %v", userID, err)
		}
	}
}

func TestRongCloud_ConversationQueryRequest(t *testing.T) {
	size, zero, negative, oversized := 100, 0, -1, 201
	for _, tt := range []struct {
		name      string
		pageToken string
		size      *int
		wantSize  string
	}{
		{name: "first page"},
		{name: "next page", pageToken: "a+&=/中文", size: &size, wantSize: "100"},
		{name: "explicit zero", size: &zero, wantSize: "0"},
		{name: "negative size forwarded", size: &negative, wantSize: "-1"},
		{name: "oversized forwarded", size: &oversized, wantSize: "201"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			client := newConversationQueryClient(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != http.MethodPost || req.URL.Path != "/user/conversation/query.json" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL)
				}
				if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
					t.Errorf("unexpected content type: %s", req.Header.Get("Content-Type"))
				}
				if req.Header.Get("User-Agent") != "rc-go-sdk/4.0.11" {
					t.Errorf("unexpected User-Agent: %s", req.Header.Get("User-Agent"))
				}
				for _, header := range []string{"App-Key", "Nonce", "Timestamp", "Signature", "X-Request-Id"} {
					if req.Header.Get(header) == "" {
						t.Errorf("missing header %s", header)
					}
				}
				if err := req.ParseForm(); err != nil {
					t.Fatal(err)
				}
				want := url.Values{"userId": {"用户+&= /"}}
				if tt.pageToken != "" {
					want.Set("pageToken", tt.pageToken)
				}
				if tt.size != nil {
					want.Set("size", tt.wantSize)
				}
				if !reflect.DeepEqual(req.PostForm, want) {
					t.Errorf("form = %v, want %v", req.PostForm, want)
				}
				return conversationQueryResponse(`{"code":200,"conversations":[{"targetId":"group1","busChannel":"channel1","conversationType":3,"version":9007199254740993}],"pageToken":"next+token"}`), nil
			})
			result, err := client.ConversationQuery(QueryConversationModel{
				UserId: "用户+&= /", PageToken: tt.pageToken, Size: tt.size,
			})
			if err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Fatalf("requests = %d, want 1", calls)
			}
			if result.PageToken != "next+token" || len(result.Conversations) != 1 {
				t.Fatalf("unexpected result: %+v", result)
			}
			item := result.Conversations[0]
			if item.TargetId != "group1" || item.BusChannel != "channel1" || item.ConversationType != GROUP || item.Version != 9007199254740993 {
				t.Errorf("unexpected conversation: %+v", item)
			}
		})
	}
}

func TestRongCloud_ConversationQueryLastPage(t *testing.T) {
	client := newConversationQueryClient(func(*http.Request) (*http.Response, error) {
		return conversationQueryResponse(`{"code":200,"conversations":[]}`), nil
	})
	result, err := client.ConversationQuery(QueryConversationModel{UserId: "user1"})
	if err != nil || result.PageToken != "" || len(result.Conversations) != 0 {
		t.Fatalf("unexpected last page: %+v, error: %v", result, err)
	}
}

func TestRongCloud_ConversationQueryErrors(t *testing.T) {
	for _, tt := range []struct {
		name string
		body string
	}{
		{"API error", `{"code":1002,"errorMessage":"invalid size"}`},
		{"malformed JSON", `{"code":200`},
		{"invalid conversation", `{"code":200,"conversations":[{"version":"invalid"}]}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			client := newConversationQueryClient(func(*http.Request) (*http.Response, error) {
				return conversationQueryResponse(tt.body), nil
			})
			_, err := client.ConversationQuery(QueryConversationModel{UserId: "user1"})
			if err == nil {
				t.Fatal("expected response error")
			}
			if tt.name == "API error" {
				if code, ok := err.(CodeResult); !ok || code.Code != 1002 || code.ErrorMessage != "invalid size" {
					t.Errorf("unexpected API error: %v", err)
				}
			}
		})
	}
	transportErr := errors.New("transport failed")
	client := newConversationQueryClient(func(*http.Request) (*http.Response, error) {
		return nil, transportErr
	})
	if _, err := client.ConversationQuery(QueryConversationModel{UserId: "user1"}); !errors.Is(err, transportErr) {
		t.Errorf("expected transport error, got %v", err)
	}
}

func TestRongCloud_ConversationTop(t *testing.T) {
	conversation := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := conversation.ConversationTop(
		PRIVATE,
		"u01",
		"u02",
		"true",
	)
	t.Log(err)
}

func TestRongCloud_ConversationMute(t *testing.T) {
	conversation := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := conversation.ConversationMute(
		PRIVATE,
		"u01",
		"u02",
	)
	t.Log(err)

}

func TestRongCloud_ConversationUnmute(t *testing.T) {
	conversation := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := conversation.ConversationMute(
		PRIVATE,
		"u01",
		"u02",
	)

	t.Log(err)

}

func TestRongCloud_ConversationGet(t *testing.T) {
	conversation := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	isMuted, err := conversation.ConversationGet(
		PRIVATE,
		"u01",
		"u02")

	t.Log(err)
	t.Log(isMuted)

}
