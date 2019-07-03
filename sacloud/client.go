package sacloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sacloud/libsacloud/v2"
)

var (
	// SakuraCloudAPIRoot APIリクエスト送信先ルートURL(末尾にスラッシュを含まない)
	SakuraCloudAPIRoot = "https://secure.sakura.ad.jp/cloud/zone"
)

// DefaultZone デフォルトゾーン、グローバルリソースなどで利用される
const DefaultZone = "is1a"

const (
	// LogLevelInfo INFOレベル
	LogLevelInfo = "INFO"
	// LogLevelWarn WARNレベル
	LogLevelWarn = "WARN"
	// LogLevelDebug DEBUGレベル
	LogLevelDebug = "DEBUG"
	// LogLevelTrace TRACEレベル
	LogLevelTrace = "TRACE"
)

// APICaller API呼び出し時に利用するトランスポートのインターフェース
type APICaller interface {
	Do(ctx context.Context, method, uri string, body interface{}) ([]byte, error)
}

// Client APIクライアント、APICallerインターフェースを実装する
type Client struct {
	// AccessToken アクセストークン
	AccessToken string `validate:"required"`
	// AccessTokenSecret アクセストークンシークレット
	AccessTokenSecret string `validate:"required"`
	// LogLevel ログレベル [TRACE / DEBUG / WARN / INFO(default)]
	LogLevel string
	// DefaultTimeoutDuration デフォルトタイムアウト間隔
	DefaultTimeoutDuration time.Duration
	// ユーザーエージェント
	UserAgent string
	// Accept-Language
	AcceptLanguage string
	// 503エラー時のリトライ回数
	RetryMax int
	// 503エラー時のリトライ待ち時間
	RetryInterval time.Duration
	// APIコール時に利用される*http.Client 未指定の場合http.DefaultClientが利用される
	HTTPClient *http.Client
}

// NewClient APIクライアント作成
func NewClient(token, tokenSecret string) *Client {
	c := &Client{
		AccessToken:            token,
		AccessTokenSecret:      tokenSecret,
		LogLevel:               LogLevelInfo,
		DefaultTimeoutDuration: 20 * time.Minute,
		UserAgent:              fmt.Sprintf("libsacloud/%s", libsacloud.Version),
		AcceptLanguage:         "",
		RetryMax:               0,
		RetryInterval:          5 * time.Second,
	}
	return c
}

// Clone APIクライアント クローン作成
func (c *Client) Clone() *Client {
	n := &Client{
		AccessToken:            c.AccessToken,
		AccessTokenSecret:      c.AccessTokenSecret,
		DefaultTimeoutDuration: c.DefaultTimeoutDuration,
		UserAgent:              c.UserAgent,
		AcceptLanguage:         c.AcceptLanguage,
		RetryMax:               c.RetryMax,
		RetryInterval:          c.RetryInterval,
		HTTPClient:             c.HTTPClient,
	}
	return n
}

func (c *Client) isOkStatus(code int) bool {
	codes := map[int]bool{
		200: true,
		201: true,
		202: true,
		204: true,
		305: false,
		400: false,
		401: false,
		403: false,
		404: false,
		405: false,
		406: false,
		408: false,
		409: false,
		411: false,
		413: false,
		415: false,
		423: false,
		500: false,
		503: false,
	}
	return codes[code]
}

// Do APIコール実施
func (c *Client) Do(ctx context.Context, method, uri string, body interface{}) ([]byte, error) {
	var (
		client = &retryableHTTPClient{
			Client:        c.HTTPClient,
			retryMax:      c.RetryMax,
			retryInterval: c.RetryInterval,
		}
		err     error
		req     *request
		strBody string
	)
	var url = uri

	if body != nil {
		var bodyJSON []byte
		bodyJSON, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		if method == "GET" {
			url = fmt.Sprintf("%s?%s", url, bytes.NewBuffer(bodyJSON))
			req, err = newRequest(ctx, method, url, nil)
		} else {
			req, err = newRequest(ctx, method, url, bytes.NewReader(bodyJSON))
		}
		b, _ := json.MarshalIndent(body, "", "\t")
		strBody = string(b)
		if c.LogLevel == LogLevelTrace {
			log.Printf("[TRACE] method : %#v , url : %s , \nbody : %s", method, url, strBody)
		}
	} else {
		req, err = newRequest(ctx, method, url, nil)
		if c.LogLevel == LogLevelTrace {
			log.Printf("[TRACE] method : %#v , url : %s ", method, url)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", url, err)
	}

	req.SetBasicAuth(c.AccessToken, c.AccessTokenSecret)
	req.Header.Add("X-Sakura-Bigint-As-Int", "1") //Use BigInt on resource ids.
	req.Header.Add("User-Agent", c.UserAgent)
	if c.AcceptLanguage != "" {
		req.Header.Add("Accept-Language", c.AcceptLanguage)
	}
	req.Method = method

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	v := &map[string]interface{}{}
	json.Unmarshal(data, v)
	b, _ := json.MarshalIndent(v, "", "\t")
	if c.LogLevel == LogLevelTrace {
		log.Printf("[TRACE] response: %s", b)
	}
	if !c.isOkStatus(resp.StatusCode) {

		errResponse := &APIErrorResponse{}
		err := json.Unmarshal(data, errResponse)

		if err != nil {
			return nil, fmt.Errorf("Error in response: %s", string(data))
		}
		return nil, NewAPIError(req.Method, req.URL, strBody, resp.StatusCode, errResponse)

	}
	if err != nil {
		return nil, err
	}

	return data, nil
}

type lenReader interface {
	Len() int
}

type request struct {
	// body is a seekable reader over the request body payload. This is
	// used to rewind the request data in between retries.
	body io.ReadSeeker

	// Embed an HTTP request directly. This makes a *Request act exactly
	// like an *http.Request so that all meta methods are supported.
	*http.Request
}

func newRequest(ctx context.Context, method, url string, body io.ReadSeeker) (*request, error) {
	var rcBody io.ReadCloser
	if body != nil {
		rcBody = ioutil.NopCloser(body)
	}

	httpReq, err := http.NewRequest(method, url, rcBody)
	if err != nil {
		return nil, err
	}

	if lr, ok := body.(lenReader); ok {
		httpReq.ContentLength = int64(lr.Len())
	}

	return &request{body, httpReq.WithContext(ctx)}, nil
}

type retryableHTTPClient struct {
	*http.Client
	retryInterval time.Duration
	retryMax      int
}

func (c *retryableHTTPClient) Do(req *request) (*http.Response, error) {
	if c.Client == nil {
		c.Client = http.DefaultClient
	}
	for i := 0; ; i++ {

		if req.body != nil {
			if _, err := req.body.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to seek body: %v", err)
			}
		}

		res, err := c.Client.Do(req.Request)
		if res != nil && res.StatusCode != http.StatusServiceUnavailable && res.StatusCode != http.StatusLocked {
			return res, err
		}
		if res != nil && res.Body != nil {
			res.Body.Close()
		}

		if err != nil {
			return res, err
		}

		remain := c.retryMax - i
		if remain == 0 {
			break
		}
		time.Sleep(c.retryInterval)
	}

	return nil, fmt.Errorf("%s %s giving up after %d attempts",
		req.Method, req.URL, c.retryMax+1)
}
