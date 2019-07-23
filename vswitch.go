package zvmsdk


import (
	"bytes"
	"encoding/json"
)


// VswitchCreateBody will be used by upper layer
// when calling vswitch create function
type VswitchCreateBody struct {
	Name string `json:"name"`
	Rdev string `json:"rdev"`
}

func getEndpointwithVswitchs(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/vswitchs")
        return buffer
}

func buildVswitchCreateRequest(body VswitchCreateBody) ([]byte) {
	data, _ := json.Marshal(body)

        return data
}

// VswitchCreate is used to create a vswitch
func VswitchCreate(endpoint string, body VswitchCreateBody) (int, []byte) {

	b := buildVswitchCreateRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	status, data := post(buffer.String(), b)

	return status, data
}

// VswitchDelete is used to delete a vswitch
func VswitchDelete(endpoint string, name string) (int, []byte) {

	buffer := getEndpointwithVswitchs(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(name)

        status, data := delete(buffer.String(), nil)

	return status, data
}

// VswitchList is used to list vswitchs, it takes no parameters
func VswitchList(endpoint string) (int, []byte) {
	buffer := getEndpointwithVswitchs(endpoint)
	status, data := get(buffer.String())

	return status, data
}

// VswitchGet is used to get vswitch info
func VswitchGet(endpoint string, name string) (int, []byte) {
        buffer := getEndpointwithVswitchs(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(name)
        status, data := get(buffer.String())

        return status, data
}
