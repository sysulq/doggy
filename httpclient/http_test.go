package doggy

import (
	"context"
	"testing"

	"net/url"
	"time"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Args    map[string]string `json:"args"`
	Headers struct {
		Accept                    string `json:"Accept"`
		Accept_Encoding           string `json:"Accept-Encoding"`
		Accept_Language           string `json:"Accept-Language"`
		Connection                string `json:"Connection"`
		Host                      string `json:"Host"`
		Upgrade_Insecure_Requests string `json:"Upgrade-Insecure-Requests"`
		User_Agent                string `json:"User-Agent"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func TestString(t *testing.T) {
	req := Get(context.Background(), "http://httpbin.org/get")
	resp, err := req.String()
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}

func TestBytes(t *testing.T) {
	req := Get(context.Background(), "http://httpbin.org/get")
	resp, err := req.Bytes()
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}

func TestToJSON(t *testing.T) {
	req := Get(context.Background(), "http://httpbin.org/get?a=b")
	foo := new(Foo)
	err := req.ToJSON(foo)
	assert.Nil(t, err)
	assert.Len(t, foo.Args, 1)
	assert.Equal(t, "http://httpbin.org/get?a=b", foo.URL)
}

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()
	req := Get(ctx, "http://google.com")
	resp, err := req.String()
	assert.Equal(t, &url.Error{
		Err: context.DeadlineExceeded,
		Op:  "Get",
		URL: "http://google.com"}, err)
	assert.Empty(t, resp)
}
