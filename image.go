package zvmsdk


import (
	"bytes"
)

// ImageCreateBody used as image create input param
type ImageCreateBody struct {
	Name string `json:"image_name"`
	RemoteHost string `json:"remote_host"`
	Meta map[string]string `json:"image_meta"`
	URL string `json:"url"`
}

// ImageUpdateBody used as image update input param
type ImageUpdateBody struct {
	DestURL string `json:"dest_url"`
	RemoteHost string `json:"remote_host"`
}


func getEndpointwithImages(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/images")
        return buffer
}

func buildImageCreateRequest(imageName string, url string, imageMeta map[string]string,
			     remoteHost string) ([]byte) {
	keys := []string{"image_name", "url", "image_meta", "remote_host"}
        values := []interface{}{imageName, url, imageMeta, remoteHost}

	return buildJSON(keys, values)
}

func buildImageUpdateRequest(destURL string, remoteHost string) ([]byte) {
        keys := []string{"dest_url", "remote_host"}
        values := []interface{}{destURL, remoteHost}

        return buildJSON(keys, values)

}

// ImageCreate creates an image
func ImageCreate(endpoint string, body ImageCreateBody) (int, []byte) {

	request := buildImageCreateRequest(body.Name, body.URL, body.Meta, body.RemoteHost)

	buffer := getEndpointwithImages(endpoint)
	status, data := post(buffer.String(), request)

	return status, data
}

// ImageDelete deletes an image
func ImageDelete(endpoint string, image string) (int, []byte) {
	buffer := getEndpointwithImages(endpoint)
	buffer.WriteString("/")
	buffer.WriteString(image)

	status, data := delete(buffer.String(), nil)

	return status, data
}

// ImageGet retrieves information from image(s)
func ImageGet(endpoint string, name string) (int, []byte) {

	buffer := getEndpointwithImages(endpoint)

	if name != "" {
		buffer.WriteString("?imagename=")
		buffer.WriteString(name)
	}
	status, data := get(buffer.String())
	return status, data
}

// ImageUpdate updates an image
func ImageUpdate(endpoint string, name string, body ImageUpdateBody) (int, []byte) {
        request := buildImageUpdateRequest(body.DestURL, body.RemoteHost)

        buffer := getEndpointwithImages(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(name)

        headers := buildAuthContext("abc")
        ctxt := RequestContext{
                values: headers,
        }

        status, data := put(buffer.String(), request, ctxt)

        return status, data

}

// ImageGetRootDiskSize gets rooot disk size of image
func ImageGetRootDiskSize(endpoint string, name string) (int, []byte) {
        buffer := getEndpointwithImages(endpoint)

	buffer.WriteString("/")
        buffer.WriteString(name)
        buffer.WriteString("/root_disk_size")

        status, data := get(buffer.String())
        return status, data

}
