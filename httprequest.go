package zvmsdk


import (
        "net/http"
	"strings"
	"io/ioutil"
)

// RequestContext used as http request context
type RequestContext struct {
	values map[string]string
}

func get(url string) (int, []byte) {
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

func post(url string, body []byte) (int, []byte) {
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

func put(url string, body []byte, context RequestContext) (int, []byte) {
        s := make([]byte, 1)

        client := &http.Client{}

        req, err := http.NewRequest("PUT", url, strings.NewReader(string(body)))

        // set content type
        req.Header.Set("Content-Type", "application/json")
        for i,v := range context.values {
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


func delete(url string, body []byte) (int, []byte){
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
