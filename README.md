# http
A HTTP Client that simplifies the usage of Go's standard net/http package.

## Features

* If you don't specify a Timeout for your request, it has a default which isn't infinite
* Easy to specify a Timeout
* Easy to send a JSON request
* Keeps Cookies across HTTP calls

## Usage

Create a new client:

    client := &http.Client{}
	
Create your own request with a specific timeout:

	req := &http.Request{
		URL:         url,
		Timeout:     10, // in seconds
		Payload:     []byte("Hello"),
		ContentType: http.ContentTypeText,
	}

Helper methods for creating common requests:

	jsonReq, err := http.NewJSONRequest(url, anyObject)

	getReq := http.NewGetRequest(url)


Send a POST request:

	resp, err := client.Post(req)

Send a GET request:

	resp, err = client.Get(req)
	
Check the Response:

    if resp.StatusCode == http.StatusOK {
		// do something useful here
	}

## Installing
Download the source of the latest version using go get:

    $ go get github.com/retrievercommunications/http

Then include in your application:

    import "github.com/retrievercommunications/http"