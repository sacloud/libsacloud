package sacloud

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"go.uber.org/ratelimit"

	"net/http"
	"sync"
)

// RateLimitRoundTripper 秒間アクセス数を制限するためのhttp.RoundTripper実装
type RateLimitRoundTripper struct {
	// Transport 親となるhttp.RoundTripper、nilの場合http.DefaultTransportが利用される
	Transport http.RoundTripper
	// RateLimitPerSec 秒あたりのリクエスト数
	RateLimitPerSec int

	once      sync.Once
	rateLimit ratelimit.Limiter
}

// RoundTrip http.RoundTripperの実装
func (r *RateLimitRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	r.once.Do(func() {
		r.rateLimit = ratelimit.New(r.RateLimitPerSec)
	})
	if r.Transport == nil {
		r.Transport = http.DefaultTransport
	}

	r.rateLimit.Take()
	return r.Transport.RoundTrip(req)
}

// TracingRoundTripper リクエスト/レスポンスのトレースログを出力するためのhttp.RoundTripper実装
type TracingRoundTripper struct {
	// Transport 親となるhttp.RoundTripper、nilの場合http.DefaultTransportが利用される
	Transport http.RoundTripper
}

// RoundTrip http.RoundTripperの実装
func (r *TracingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if r.Transport == nil {
		r.Transport = http.DefaultTransport
	}

	reqView := struct {
		Method string
		URL    string
		Header http.Header
		Body   string
	}{
		Method: req.Method,
		URL:    req.URL.String(),
		Header: req.Header,
	}
	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		reqView.Body = string(body)
	}
	if data, err := json.Marshal(reqView); err == nil {
		log.Printf("[TRACE] \trequest: %s", string(data))
	}

	res, err := r.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if res != nil {
		resView := struct {
			Status     string
			StatusCode int
			Header     http.Header
			Body       string
		}{
			Status:     res.Status,
			StatusCode: res.StatusCode,
			Header:     res.Header,
		}
		if res.Body != nil {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			resView.Body = string(body)
			res.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		if data, err := json.Marshal(resView); err == nil {
			log.Printf("[TRACE] \tresponse: %s", string(data))
		}
	}

	return res, err
}
