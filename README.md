# http
A HTTP Client that simplifies the usage of Go's standard net/http package.

So why not use the standard package directly?
Well you can but this just simplifies things if you are a Go newbies or have a simple use case.

## Features

* If you don't specify a Timeout for your request, it has a default which isn't infinite
* Easy to specify a Timeout
* Easy to send a JSON request
* Keeps Cookies across HTTP calls

## Usage

Create a new client:

    client := &http.Client{}
	
Create your own request with a specific timeout:

	req := &Request{
		URL:         url,
		Timeout:     10, // in seconds
		Payload:     []byte("Hello"),
		ContentType: ContentTypeText,
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
Using this client is easy. Just use go get to install the latest version
of the library.

    $ go get github.com/grange74/http

Next include in your application.

    import "github.com/grange74/http"