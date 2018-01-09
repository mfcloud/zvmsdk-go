package zvmsdk


import (
	"bytes"
)

type VswitchCreateBody struct {
	Name string `json:"name"`
	Rdev string `json:"rdev"`
}


type VswitchGrantBody struct {
	GrantUserId string `json:"grant_userid"`
}

type VswitchRevokeBody struct {
	RevokeUserId string `json:"revoke_userid"`
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

	return buildJson(keys, values)
}

func VswitchCreate(endpoint string, body VswitchCreateBody) (int, []byte) {

	bodyJson := buildVswitchCreateRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	status, data := post(buffer.String(), bodyJson)

	return status, data
}

func VswitchDelete(endpoint string, name string) (int, []byte) {

	buffer := getEndpointwithVswitchs(endpoint)
	buffer.WriteString("/")
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

func buildVswitchGrantRequest(body VswitchGrantBody) ([]byte) {
        keys := []string{"grant_userid"}
        values := []interface{}{body.GrantUserId}

	return buildJson(keys, values)
}

func buildVswitchRevokeRequest(body VswitchRevokeBody) ([]byte) {
        keys := []string{"revoke_userid"}
        values := []interface{}{body.RevokeUserId}

	return buildJson(keys, values)
}


func VswitchGrant(endpoint string, name string, body VswitchGrantBody) (int, []byte) {
	bodyJson := buildVswitchGrantRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	buffer.WriteString("/")
	buffer.WriteString(name)

	headers := buildAuthContext("abc")
	ctxt := RequestContext{
		values: headers,
	}

        status, data := put(buffer.String(), bodyJson, ctxt)

	return status, data
}


func VswitchRevoke(endpoint string, name string, body VswitchRevokeBody) (int, []byte) {
        bodyJson := buildVswitchRevokeRequest(body)

        buffer := getEndpointwithVswitchs(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(name)

        headers := buildAuthContext("bcd")
        ctxt := RequestContext{
                                values: headers,
        }


        status, data := put(buffer.String(), bodyJson, ctxt)

        return status, data
}

func VswitchGet(endpoint string, name string) (int, []byte) {
        buffer := getEndpointwithVswitchs(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(name)
        status, data := get(buffer.String())

        return status, data
}
