package zvmsdk


import (
	"bytes"
)


// VswitchCreateBody will be used by upper layer
// when calling vswitch create function
type VswitchCreateBody struct {
	Name string `json:"name"`
	Rdev string `json:"rdev"`
}

// VswitchGrantBody will be used by upper layer
// when calling vswitch grant function
type VswitchGrantBody struct {
	GrantUserID string `json:"grant_userid"`
}

// VswitchRevokeBody will be used by upper layer
// when calling vswitch revoke function
type VswitchRevokeBody struct {
	RevokeUserID string `json:"revoke_userid"`
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

	return buildJSON(keys, values)
}

// VswitchCreate is used to create a vswitch
func VswitchCreate(endpoint string, body VswitchCreateBody) (int, []byte) {

	bodyJSON := buildVswitchCreateRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	status, data := post(buffer.String(), bodyJSON)

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

func buildVswitchGrantRequest(body VswitchGrantBody) ([]byte) {
        keys := []string{"grant_userid"}
        values := []interface{}{body.GrantUserID}

	return buildJSON(keys, values)
}

func buildVswitchRevokeRequest(body VswitchRevokeBody) ([]byte) {
        keys := []string{"revoke_userid"}
        values := []interface{}{body.RevokeUserID}

	return buildJSON(keys, values)
}

// VswitchGrant is used to grant guest to vswitch
func VswitchGrant(endpoint string, name string, body VswitchGrantBody) (int, []byte) {
	bodyJSON := buildVswitchGrantRequest(body)

	buffer := getEndpointwithVswitchs(endpoint)
	buffer.WriteString("/")
	buffer.WriteString(name)

	headers := buildAuthContext("abc")
	ctxt := RequestContext{
		values: headers,
	}

        status, data := put(buffer.String(), bodyJSON, ctxt)

	return status, data
}

// VswitchRevoke is used to revoke guest from vswitch
func VswitchRevoke(endpoint string, name string, body VswitchRevokeBody) (int, []byte) {
        bodyJSON := buildVswitchRevokeRequest(body)

        buffer := getEndpointwithVswitchs(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(name)

        headers := buildAuthContext("bcd")
        ctxt := RequestContext{
                                values: headers,
        }


        status, data := put(buffer.String(), bodyJSON, ctxt)

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
