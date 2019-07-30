package zvmsdk

import (
	"bytes"
	"encoding/json"
)

// ImageCreateBody used as image create input param
type ImageCreateBody struct {
	Name       string            `json:"image_name"`
	RemoteHost string            `json:"remote_host,omitempty"`
	Meta       map[string]string `json:"image_meta,omitempty"`
	URL        string            `json:"url"`
}

// ImageUpdateBody used as image update input param
type ImageUpdateBody struct {
	DestURL    string `json:"dest_url"`
	RemoteHost string `json:"remote_host,omitempty"`
}

func getEndpointwithImages(endpoint string) bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(endpoint)
	buffer.WriteString("/images")
	return buffer
}

func buildImageCreateRequest(body ImageCreateBody) []byte {
	data, _ := json.Marshal(body)

	return data
}

func buildImageUpdateRequest(body ImageUpdateBody) []byte {
	data, _ := json.Marshal(body)

	return data
}

// ImageCreate creates an image
func ImageCreate(endpoint string, body ImageCreateBody) (int, []byte) {

	request := buildImageCreateRequest(body)

	buffer := getEndpointwithImages(endpoint)
	status, data := hq.Post(buffer.String(), request)

	return status, data
}

// ImageDelete deletes an image
func ImageDelete(endpoint string, image string) (int, []byte) {
	buffer := getEndpointwithImages(endpoint)
	buffer.WriteString("/")
	buffer.WriteString(image)

	status, data := hq.Delete(buffer.String(), nil)

	return status, data
}

// ImageGet retrieves information from image(s)
func ImageGet(endpoint string, name string) (int, []byte) {

	buffer := getEndpointwithImages(endpoint)

	if name != "" {
		buffer.WriteString("?imagename=")
		buffer.WriteString(name)
	}
	status, data := hq.Get(buffer.String())
	return status, data
}

// ImageUpdate updates an image
func ImageUpdate(endpoint string, name string, body ImageUpdateBody) (int, []byte) {
	request := buildImageUpdateRequest(body)

	buffer := getEndpointwithImages(endpoint)
	buffer.WriteString("/")
	buffer.WriteString(name)

	headers := buildAuthContext("abc")
	ctxt := RequestContext{
		values: headers,
	}

	status, data := hq.Put(buffer.String(), request, ctxt)

	return status, data

}

// ImageGetRootDiskSize gets rooot disk size of image
func ImageGetRootDiskSize(endpoint string, name string) (int, []byte) {
	buffer := getEndpointwithImages(endpoint)

	buffer.WriteString("/")
	buffer.WriteString(name)
	buffer.WriteString("/root_disk_size")

	status, data := hq.Get(buffer.String())
	return status, data

}
