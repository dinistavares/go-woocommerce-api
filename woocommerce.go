package woocommerce

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultRestEndpointVersion   = "v3"
	defaultHeaderName            = "Authorization"
	acceptedContentType          = "application/json"
	userAgent                    = "go-woocommerce-api/1.1"
	clientRequestRetryAttempts   = 2
	clientRequestRetryHoldMillis = 1000
)

var errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
var errorDoAttemptNilRequest = errors.New("request could not be constructed")

type ClientConfig struct {
	HttpClient          *http.Client
	RestEndpointURL     string
	RestEndpointVersion string
}

type auth struct {
	HeaderName string
	ApiKey     string
}

type Client struct {
	config  *ClientConfig
	client  *http.Client
	auth    *auth
	baseURL *url.URL

	Coupons           *CouponsService
	Customers         *CustomersService
	Orders            *OrdersService
	OrderNotes        *OrderNotesService
	Refunds           *RefundsService
	Products          *ProductsService
	ProductTags       *ProductTagService
	ProductVariations *ProductVariationService
	Webhooks          *WebhookService
}

type service struct {
	client *Client
}

type errorResponse struct {
	Response *http.Response

	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    ErrorData `json:"data"`
}

type ErrorData struct {
	Status int `json:"status"`
}

func (response *errorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		response.Response.Request.Method, response.Response.Request.URL,
		response.Response.StatusCode, response.Message)
}

func New(shopURL string) (*Client, error) {
	if shopURL == "" {
		return nil, errors.New("store url is required")
	}

	httpClient := http.DefaultClient

	if shopURL == "https://wordpress.local.sparklayer.io" {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	config := ClientConfig{
		HttpClient:          http.DefaultClient,
		RestEndpointVersion: defaultRestEndpointVersion,
		RestEndpointURL:     shopURL,
	}

	//config.HttpClient = http.DefaultClient
	//config.RestEndpointURL = shopURL
	//config.RestEndpointVersion = defaultRestEndpointVersion

	// Create client
	baseURL, err := url.Parse(config.RestEndpointURL + "/wp-json/wc/" + defaultRestEndpointVersion)

	if err != nil {
		return nil, err
	}

	client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL}

	// Map services
	client.Coupons = &CouponsService{client: client}
	client.Customers = &CustomersService{client: client}
	client.Orders = &OrdersService{client: client}
	client.OrderNotes = &OrderNotesService{client: client}
	client.Refunds = &RefundsService{client: client}
	client.Products = &ProductsService{client: client}
	client.ProductTags = &ProductTagService{client: client}
	client.ProductVariations = &ProductVariationService{client: client}
	client.Webhooks = &WebhookService{client: client}

	return client, nil
}

// Authenticate saves authentication parameters for user
func (client *Client) Authenticate(consumer_key string, consumer_secret string) {
	client.auth.HeaderName = defaultHeaderName

	client.auth.ApiKey = "Basic " + base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{consumer_key, consumer_secret}, ":")))
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, opts interface{}, body interface{}) (*http.Request, error) {
	// Append Query Params to URL
	if opts != nil {
		queryParams, err := query.Values(opts)

		if err != nil {
			return nil, err
		}

		rawQuery := queryParams.Encode()

		if rawQuery != "" {
			urlStr += "?" + rawQuery
		}
	}

	rel, err := url.Parse(client.config.RestEndpointVersion + urlStr)
	if err != nil {
		return nil, err
	}

	reqUrl := client.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, reqUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add(client.auth.HeaderName, client.auth.ApiKey)
	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("Content-type", acceptedContentType)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	var lastErr error

	attempts := 0

	for attempts < clientRequestRetryAttempts {
		// Hold before this attempt? (i.e. not first attempt)
		if attempts > 0 {
			time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
		}

		// Dispatch request attempt
		attempts++
		resp, shouldRetry, err := client.doAttempt(req, v)

		// Return response straight away? (we are done)
		if !shouldRetry {
			return resp, err
		}

		// Should retry: store last error (we are not done)
		lastErr = err
	}

	// Set default error? (all attempts failed, but no error is set)
	if lastErr == nil {
		lastErr = errorDoAllAttemptsExhausted
	}

	// All attempts failed, return last attempt error
	return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*http.Response, bool, error) {
	if req == nil {
		return nil, false, errorDoAttemptNilRequest
	}

	resp, err := client.client.Do(req)

	if checkRequestRetry(resp, err) {
		return nil, true, err
	}

	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, false, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return resp, false, err
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
	// Low-level error, or response status is a server error? (HTTP 5xx)
	if err != nil || response.StatusCode >= 500 {
		return true
	}

	// No low-level error (should not retry)
	return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &errorResponse{Response: response}

	data, err := io.ReadAll(response.Body)
	if err == nil && data != nil {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}
