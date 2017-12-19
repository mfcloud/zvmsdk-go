package zvmsdk


import (
	"fmt"
	"bytes"
	"encoding/json"
)

type ImageCreateBody struct {
	Image_name string `json:"image_name"`
	Remote_host string `json:"remote_host"`
	Image_meta map[string]string `json:"image_meta"`
	Url string `json:"url"`
}


func buildImageCreateRequest(image_name string, url string, image_meta map[string]string,
			     remote_host string) ([]byte) {
	keys := []string{"image_name", "url", "image_meta", "remote_host"}
        values := []interface{}{image_name, url, image_meta, remote_host}

        // map values to keys
        m := make(map[string]interface{})
        for i,v := range values {
                m[keys[i]] = v
        }
        // convert map to JSON
        data, _ := json.Marshal(m)

	return data
}

func ImageCreate(image_name string, url string, image_meta map[string]string,
		 remote_host string) {

	data := buildImageCreateRequest(image_name, url, image_meta, remote_host)

	res, result := post("http://localhost:8080", data)
	fmt.Println("output is ", res, string(result))
}

func ImageDelete() {
	res, result := delete("http://localhost:8080", nil)
	fmt.Println("output is ", res, string(result))
}

func ImageGet(name string) {
	var buffer bytes.Buffer

	buffer.WriteString("http://localhost:8080/images")

	if name != "" {
		buffer.WriteString("?imagename=")
		buffer.WriteString(name)
	}
	res, result := get(buffer.String())

	fmt.Println("output is ", res, string(result))
}


