package zvmsdk


import (
	"fmt"
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

func VswitchCreate(body VswitchCreateBody) {

	data := buildVswitchCreateRequest(body)

	res, result := post("http://localhost:8080", data)
	fmt.Println("output is ", res, string(result))
}

func VswitchDelete(name string) {
	res, result := delete("http://localhost:8080", nil)
	fmt.Println("output is ", res, string(result))
}

//vswitch list takes no param
func VswitchList() {
	var buffer bytes.Buffer

	buffer.WriteString("http://localhost:8080/vswitchs")

	res, result := get(buffer.String())

	fmt.Println("output is ", res, string(result))
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

func VswitchUpdate(name string, body VswitchUpdateBody) {
	data := buildVswitchUpdateRequest(body)

	var buffer bytes.Buffer

        buffer.WriteString("http://localhost:8080/vswitchs/")
	buffer.WriteString(name)
        res, result := put(buffer.String(), data)

        fmt.Println("output is ", res, string(result))
}
