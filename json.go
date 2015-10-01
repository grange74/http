package http

import "encoding/json"

// NewJSONRequest creates a new HTTP Request with a JSON Payload
func NewJSONRequest(url string, jsonObj interface{}) (*Request, error) {

	b, err := json.Marshal(jsonObj)

	if err != nil {
		return nil, err
	}

	req := &Request{
		URL:         url,
		Timeout:     DefaultTimeout, // default in seconds
		Payload:     b,
		ContentType: ContentTypeJSON,
	}

	return req, nil
}
