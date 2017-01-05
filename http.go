package doggy

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/facebookgo/httpcontrol"
	"github.com/hnlq715/doggy/middleware"
	"github.com/uber-go/zap"
)

var (
	defaultClient *http.Client
)

func initHttp() {
	defaultClient = &http.Client{
		Transport: &httpcontrol.Transport{
			RequestTimeout: config.HttpClient.Timeout,
			MaxTries:       config.HttpClient.Retry,
		},
	}
}

func ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(config.Listen, handler)
}

func newRequest(ctx context.Context, method, url string, body []byte) ([]byte, error) {
	l := middleware.LogFromContext(ctx).With(zap.String("query", url), zap.String("type", "http"), zap.String("direction", "out"))
	now := time.Now()

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		l.Error("http.NewRequest failed", zap.Error(err))
		return nil, err
	}

	if method == "POST" && body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := defaultClient.Do(req)
	if err != nil {
		l.Error("defaultClient.Get failed", zap.Error(err), zap.Float64("responsetime", time.Now().Sub(now).Seconds()))
		return nil, err
	}
	defer resp.Body.Close()

	var data []byte
	l = l.With(zap.Float64("responsetime", time.Now().Sub(now).Seconds()))

	if resp.StatusCode != 200 {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			l.Error("ioutil.ReadAll failed", zap.Error(err))
			return nil, err
		}
		return nil, fmt.Errorf("statusCode:%d not expected", resp.StatusCode)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Error("ioutil.ReadAll failed", zap.Error(err))
		return nil, err
	}

	l.Info("newRequest finished")

	return data, nil
}

func Get(ctx context.Context, url string) ([]byte, error) {
	return newRequest(ctx, "GET", url, nil)
}

func Post(ctx context.Context, url string, body []byte) ([]byte, error) {
	return newRequest(ctx, "POST", url, body)
}
