package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {}


func (c *Client) Post(url string, body io.Reader) ([]byte, error) {
	return request(url, "POST", body)
}

func (c *Client) Get(url string) ([]byte, error) {
	return request(url, "GET", nil)
}

func request(url string, method string, body io.Reader) ([]byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest(method, url, body)

	if method == "POST" {
		request.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)

	if response.StatusCode == http.StatusOK {
		return buf, nil
	} else {
		return nil, errors.New(fmt.Sprint("code=", response.StatusCode, " message=", string(buf)))
	}
}