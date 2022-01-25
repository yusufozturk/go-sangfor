package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MockTransport struct {
	Response MockResponse
	Error    error
}

func NewMockTransport(statusCode int, headers map[string]string, body map[string][]byte) *MockTransport {
	return &MockTransport{
		Response: NewMockResponse(statusCode, headers, body),
	}
}

// RoundTrip receives HTTP requests and routes them to the appropriate
// responder.  It is required to implement the http.RoundTripper
// interface.  You will not interact with this directly, instead the
// *http.Client you are using will call it for you.
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	return m.Response.MakeResponse(req), nil
}

type MockResponse struct {
	StatusCode int
	HeadersMap map[string]string
	Body       map[string][]byte
}

func NewMockResponse(statusCode int, headers map[string]string, body map[string][]byte) MockResponse {
	if headers == nil {
		headers = map[string]string{}
	}

	if body == nil {
		body = map[string][]byte{}
	}

	return MockResponse{StatusCode: statusCode, HeadersMap: headers, Body: body}
}

func (r *MockResponse) MakeResponse(req *http.Request) *http.Response {
	// Get Status
	status := strconv.Itoa(r.StatusCode) + " " + http.StatusText(r.StatusCode)

	// Get HTTP Headers
	header := http.Header{}
	for name, value := range r.HeadersMap {
		header.Set(name, value)
	}

	// Get HTTP Response
	httpResponse := r.Body[req.URL.String()]
	if httpResponse == nil {
		// Print Empty Request
		fmt.Println(req.URL.String())

		// Get First Available Process
		for _, response := range r.Body {
			httpResponse = response
			break
		}
	}

	// Set Content Length
	contentLength := len(httpResponse)
	header.Set("Content-Length", strconv.Itoa(contentLength))

	// Set HTTP Response
	res := &http.Response{
		Status:           status,
		StatusCode:       r.StatusCode,
		Proto:            "HTTP/1.0",
		ProtoMajor:       1,
		ProtoMinor:       0,
		Header:           header,
		Body:             ioutil.NopCloser(bytes.NewReader(httpResponse)),
		ContentLength:    int64(contentLength),
		TransferEncoding: []string{},
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          req,
		TLS:              nil,
	}

	// Should no set Content-Length header when 204 or 304
	if r.StatusCode == http.StatusNoContent || r.StatusCode == http.StatusNotModified {
		if res.ContentLength != 0 {
			res.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))
			res.ContentLength = 0
		}
		header.Del("Content-Length")
	}

	return res
}

type MockDatabase struct {
	Database map[string][]byte
}

func (d *MockDatabase) AddToDatabase(key string, value interface{}) {
	// Convert to JSON
	jsonResponse, err := json.Marshal(value)
	if err != nil {
		return
	}

	// Check Database
	if d.Database == nil {
		d.Database = make(map[string][]byte)
	}

	// Add to Database
	d.Database[key] = jsonResponse
}
