package zvmsdk


import (
        "net/http"
	"io/ioutil"
)

type Body []byte


func get(url string) ([]byte) {
        var s []byte = make([]byte, 1)

	client := &http.Client{}

        req, err := http.NewRequest("GET", url, nil)

        // set content type
        req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return s
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s
	}

        return result
}

