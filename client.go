// HTTP Client that simplifies HTTP interactions in Go by
// wrapping the built-in net/http package.

package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Client represents a HTTP client that sends Requests and receives Responses.
type Client struct {
	cookies []*http.Cookie // maintained for the entire client session
}

// Request represents an HTTP request sent by a client.
type Request struct {
	URL         string
	Timeout     int // in seconds
	Payload     []byte
	ContentType string
}

// Response represents the response from an HTTP request.
type Response struct {
	StatusCode int // e.g. 200
	Payload    []byte
}

// NewRequest returns a new Request given a URL, payload and contentType.
func NewRequest(url, payload, contentType string) *Request {
	return &Request{
		URL:         url,
		Timeout:     DefaultTimeout, // default in seconds
		Payload:     []byte(payload),
		ContentType: contentType,
	}
}

// NewGetRequest returns a new Request given just a URL.
func NewGetRequest(url string) *Request {
	return &Request{
		URL:         url,
		Timeout:     DefaultTimeout, // default in seconds
		Payload:     nil,
		ContentType: ContentTypeNone,
	}
}

// Get sends a HTTP GET request to a server
func (c *Client) Get(req *Request) (*Response, error) {
	return c.Send(req, MethodGet)
}

// Post sends a HTTP Post request to a server
func (c *Client) Post(req *Request) (*Response, error) {
	return c.Send(req, MethodPost)
}

// Send a HTTP request of any method type to a server
func (c *Client) Send(request *Request, method string) (*Response, error) {

	timeout := time.Duration(time.Duration(request.Timeout) * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, request.URL, bytes.NewReader(request.Payload))

	// Not all requests require a ContentType (e.g. GETs)
	if request.ContentType != ContentTypeNone {
		req.Header.Set("Content-Type", request.ContentType)
	}

	// Add any cookies that have been obtained from
	// any previous calls on this client.
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}

	// make the actual HTTP request
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// Keep the Cookies for the next request
	for _, cookie := range resp.Cookies() {
		if index := getCookieIndex(cookie.Name, c.cookies); index > -1 {
			// update existing cookie
			c.cookies[index] = cookie
		} else {
			// add new cookies
			c.cookies = append(c.cookies, cookie)
		}
	}

	response := &Response{
		StatusCode: resp.StatusCode,
		Payload:    payload,
	}

	return response, nil
}

func getCookieIndex(name string, cookies []*http.Cookie) int {
	for index, cookie := range cookies {
		if cookie.Name == name {
			return index
		}
	}
	return -1
}
