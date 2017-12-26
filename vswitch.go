package zvmsdk


import (
	"bytes"
	"encoding/json"
)

type VswitchCreateBody struct {
	Name string `json:"name"`
	Rdev string `json:"rdev"`
}


type VswitchUpdateBody struct {
	GrantUserId string `json:"grant_userid"`
}

func getEndpointwithVswitchs(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/vswitchs")
        return buffer
}


func buildVswitchCreateRequest(body VswitchCreateBody) ([]byte) {
	keys := []string{"name", "rdev"}
        values := []interface{}{body.Name, body.Rdev}

        // map values to keys
        m := make(map[string]interface{})
        for i,v := range values {
                m[keys[i]] = v
        }
        // convert map to JSON
        data, _ := json.Marshal(m)

	return data
}

func VswitchCreate(endpoint string, body VswitchCreateBody) (int, []byte) {

	bodyJson := buildVswitchCreateRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	status, data := post(buffer.String(), bodyJson)

	return status, data
}

func VswitchDelete(endpoint string, name string) (int, []byte) {

	buffer := getEndpointwithVswitchs(endpoint)

        buffer.WriteString(name)
        status, data := delete(buffer.String(), nil)

	return status, data
}

//vswitch list takes no param
func VswitchList(endpoint string) (int, []byte) {
	buffer := getEndpointwithVswitchs(endpoint)
	status, data := get(buffer.String())

	return status, data
}

func buildVswitchUpdateRequest(body VswitchUpdateBody) ([]byte) {
        keys := []string{"grant_userid"}
        values := []interface{}{body.GrantUserId}

        // map values to keys
        m := make(map[string]interface{})
        for i,v := range values {
                m[keys[i]] = v
        }
        // convert map to JSON
        data, _ := json.Marshal(m)

        return data
}

func VswitchUpdate(endpoint string, name string, body VswitchUpdateBody) (int, []byte) {
	bodyJson := buildVswitchUpdateRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	buffer.WriteString(name)
        status, data := put(buffer.String(), bodyJson)

	return status, data
}
