package zvmsdk

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// RequestContext used as http request context
type RequestContext struct {
	values map[string]string
}

type IHttprequest interface {
	Get(url string) (int, []byte)
	Post(url string, body []byte) (int, []byte)
	Put(url string, body []byte, context RequestContext) (int, []byte)
	Delete(url string, body []byte) (int, []byte)
}

type HttpReq struct {
}

func (h *HttpReq) Get(url string) (int, []byte) {
	s := make([]byte, 1)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)

	if err != nil {
		return -1, s
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, s
	}

	return resp.StatusCode, result
}

func (h *HttpReq) Post(url string, body []byte) (int, []byte) {
	s := make([]byte, 1)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))

	// set content type
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return -1, s
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, s
	}

	return resp.StatusCode, result
}

func (h *HttpReq) Put(url string, body []byte, context RequestContext) (int, []byte) {
	s := make([]byte, 1)

	client := &http.Client{}

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(body)))

	// set content type
	req.Header.Set("Content-Type", "application/json")
	for i, v := range context.values {
		req.Header.Set(i, v)
	}

	resp, err := client.Do(req)

	if err != nil {
		return -1, s
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, s
	}

	return resp.StatusCode, result
}

func (h *HttpReq) Delete(url string, body []byte) (int, []byte) {
	s := make([]byte, 1)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", url, nil)

	// set content type
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return -1, s
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, s
	}

	return resp.StatusCode, result
}
