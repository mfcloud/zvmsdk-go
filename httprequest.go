package zvmsdk


import (
        "net/http"
	"io/ioutil"
)



func get(url string) (int, []byte) {
        var s []byte = make([]byte, 1)

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
        var s []byte = make([]byte, 1)

	client := &http.Client{}

        req, err := http.NewRequest("POST", url, nil)

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

func put(url string, body []byte) (int, []byte) {
        var s []byte = make([]byte, 1)

        client := &http.Client{}

        req, err := http.NewRequest("PUT", url, nil)

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


func delete(url string, body []byte) (int, []byte){
        var s []byte = make([]byte, 1)

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
