package ki

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func (client *Client) Init() *Client {
	client.httpClient = &http.Client{}

	return client
}

func (client *Client) HttpGet(url string) (data []byte, err error) {
	var (
		res *http.Response
	)

	res, err = http.Get(url)
	if err != nil {
		return
	}

	if res.Body == nil {
		return
	}

	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) HttpGetJson(url string, dst any) (err error) {
	var (
		data []byte
	)

	data, err = client.HttpGet(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, dst)
	if err != nil {
		return
	}

	return
}

func (client *Client) HttpPost(url string, data []byte) (result []byte, err error) {
	var (
		req *http.Request
		res *http.Response
	)

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}

	req.Header = http.Header{}

	res, err = client.httpClient.Do(req)
	if err != nil {
		return
	}

	if res.Body == nil {
		return
	}

	defer res.Body.Close()

	result, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	return
}

func (client *Client) HttpPostJson(url string, src any, dst any) (err error) {
	var (
		data []byte
	)

	data, err = json.Marshal(src)
	if err != nil {
		return
	}

	data, err = client.HttpPost(url, data)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, dst)
	if err != nil {
		return
	}

	return
}
