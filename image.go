package zvmsdk


import (
	"bytes"
)

type ImageCreateBody struct {
	Name string `json:"image_name"`
	RemoteHost string `json:"remote_host"`
	Meta map[string]string `json:"image_meta"`
	Url string `json:"url"`
}

type ImageUpdateBody struct {
	DestUrl string `json:"dest_url"`
	RemoteHost string `json:"remote_host"`
}


func getEndpointwithImages(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/images")
        return buffer
}

func buildImageCreateRequest(image_name string, url string, image_meta map[string]string,
			     remote_host string) ([]byte) {
	keys := []string{"image_name", "url", "image_meta", "remote_host"}
        values := []interface{}{image_name, url, image_meta, remote_host}

	return buildJson(keys, values)
}

func buildImageUpdateRequest(dest_url string, remote_host string) ([]byte) {
        keys := []string{"dest_url", "remote_host"}
        values := []interface{}{dest_url, remote_host}

        return buildJson(keys, values)

}

func ImageCreate(endpoint string, body ImageCreateBody) (int, []byte) {

	request := buildImageCreateRequest(body.Name, body.Url, body.Meta, body.RemoteHost)

	buffer := getEndpointwithImages(endpoint)
	status, data := post(buffer.String(), request)

	return status, data
}

func ImageDelete(endpoint string, image string) (int, []byte) {
	buffer := getEndpointwithImages(endpoint)
	buffer.WriteString("/")
	buffer.WriteString(image)

	status, data := delete(buffer.String(), nil)

	return status, data
}

func ImageGet(endpoint string, name string) (int, []byte) {

	buffer := getEndpointwithImages(endpoint)

	if name != "" {
		buffer.WriteString("?imagename=")
		buffer.WriteString(name)
	}
	status, data := get(buffer.String())
	return status, data
}


func ImageUpdate(endpoint string, name string, body ImageUpdateBody) (int, []byte) {
        request := buildImageUpdateRequest(body.DestUrl, body.RemoteHost)

        buffer := getEndpointwithImages(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(name)
        status, data := put(buffer.String(), request)

        return status, data

}

func ImageGetRootDiskSize(endpoint string, name string) (int, []byte) {
        buffer := getEndpointwithImages(endpoint)

	buffer.WriteString("/")
        buffer.WriteString(name)
        buffer.WriteString("/root_disk_size")

        status, data := get(buffer.String())
        return status, data

}
