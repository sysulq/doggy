package doggy

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/uber-go/zap"
)

func ListenAndServe(handler http.Handler) error {
	return http.ListenAndServe(config.Listen, handler)
}

func NewRequest(ctx context.Context, method, url string, body []byte) ([]byte, error) {
	l := LogFromContext(ctx).With(zap.String("query", url), zap.String("type", "http"), zap.String("direction", "out"))
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		l.Error("http.NewRequest failed", zap.Error(err))
		return nil, err
	}

	if method == "POST" && body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	now := time.Now()
	http.DefaultClient.Timeout = config.HttpClient.Timeout

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l.Error("http.DefaultClient.Do failed", zap.Error(err), zap.Float64("responsetime", time.Now().Sub(now).Seconds()))
		return nil, err
	}

	var data []byte
	l = l.With(zap.Float64("responsetime", time.Now().Sub(now).Seconds()))

	defer resp.Body.Close()

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

	l.Info("newRequest finished", zap.String("POST", string(body)))

	return data, nil
}
