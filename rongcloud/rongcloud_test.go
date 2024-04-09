package rongcloud

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestNewRongCloud(t *testing.T) {
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))
	dialer := &net.Dialer{
		Timeout:   123 * time.Second,
		KeepAlive: 123 * time.Second,
	}
	globalTransport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			log.Printf("custom dialer works, timeout %s keepalive %s", dialer.Timeout, dialer.KeepAlive)
			return dialer.DialContext(ctx, network, addr)
		},
		MaxIdleConnsPerHost: 123,
	}
	ctx := context.Background()
	_, err := rc.UserCheckOnline(ctx, &UserCheckOnlineRequest{UserId: StringPtr("u01")})
	if err != nil {
		log.Fatalf("usercheck online error %s", err)
	}

	oldTransport, ok := rc.httpClient.Transport.(*http.Transport)
	if !ok {
		log.Fatalf("RoundTrip convert to Transport failed")
	}
	rc.SetHttpTransport(globalTransport)
	newTransport, ok := rc.GetHttpTransport().(*http.Transport)
	if !ok {
		log.Fatalf("RoundTrip convert to Transport failed")
	}
	log.Printf("oldTransport MaxIdleConnsPerHost %d, newTransport MaxIdleConnsPerHost %d", oldTransport.MaxIdleConnsPerHost, newTransport.MaxIdleConnsPerHost)
	_, err = rc.UserCheckOnline(ctx, &UserCheckOnlineRequest{
		UserId: StringPtr("u01"),
	})
	if err != nil {
		log.Fatalf("user check online 2 error %s", err)
	}
}

func TestNewRongCloudInit(t *testing.T) {
	dialer := &net.Dialer{
		Timeout:   123 * time.Second,
		KeepAlive: 123 * time.Second,
	}
	globalTransport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			log.Printf("custom dialer works, timeout %s keepalive %s", dialer.Timeout, dialer.KeepAlive)
			return dialer.DialContext(ctx, network, addr)
		},
	}
	rc := NewRongCloud(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"), WithTransport(globalTransport), WithTimeout(1*time.Hour), WithKeepAlive(1*time.Hour)) // transport优先级高于其他http参数
	ctx := context.Background()
	_, err := rc.UserCheckOnline(ctx, &UserCheckOnlineRequest{UserId: StringPtr("u01")})
	if err != nil {
		log.Fatalf("usercheck online error %s", err)
	}
}
