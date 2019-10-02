package transport

import (
	"io/ioutil"
	"errors"
	"net/http"
	"net/url"
	"path"
)

// HTTPTransport is a transport via HTTP protocol
type HTTPTransport struct {
	URL string
}

// NewHTTPTransport creates HTTPTransport instance
func NewHTTPTransport(url string) HTTPTransport {
	return HTTPTransport{
		URL: url,
	}
}

// Send sends a message via HTTP
func (t *HTTPTransport) Send(m Message) (bool, error) {
	return false, errors.New("httptransport: Send not implemented")
}

// Receive a message via HTTP
func (t *HTTPTransport) Receive(rp RequestParams) (chan Message, chan error) {
	messageChannel := make(chan Message, 1)
	defer close(messageChannel)
	errorChannel := make(chan error, 1)
	defer close(errorChannel)

	go func(){
		// Creating proper URL
		u, err := url.Parse(t.URL)
		if err != nil {
			errorChannel <- err
			return
		}
		u.Path = path.Join(u.Path, rp["HTTPPath"])

		// Creating a request
		req, err := http.NewRequest(rp["HTTPMethod"], u.String(), nil)
		if err != nil {
			errorChannel <- err
			return
		}

		// Extracting data
		data := make(map[string]string)
		for k, v := range map[string]string(rp) {
			if k[:4] != "HTTP" {
				data[k] = v
			}
		}

		// Making a request
		var resp *http.Response
		switch rp["HTTPMethod"] {
		case "GET":
			q := u.Query()
			for k, v := range data {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
			resp, err = http.Get(u.String())
		default:
			formData := url.Values{}
			for k, v := range data {
				formData.Add(k, v)
			}
			resp, err = http.PostForm(u.String(), formData)
		}
		if err != nil {
			errorChannel <- err
			return
		}
		
		// Processing results
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errorChannel <- err
			return
		}
		messageChannel <- Message{
			Body: body,
		}
	}()

	return messageChannel, errorChannel
}

