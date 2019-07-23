package zvmsdk


import (
	"bytes"
	"encoding/json"
)


type coupleInfo struct {
	Couple int `json:"couple"`
	Active int `json:"active"`
	Vswitch string `json:"vswitch"`
}

type coupleBody struct {
	Info   coupleInfo `json:"info,omitempty"`
}

type grantVswitch struct {
	Userid string `json:"grant_userid,omitempty"`
}

type grantBody struct {
	Vswitch grantVswitch `json:"vswitch,omitempty"`
}

// VswitchCreateBody will be used by upper layer
// when calling vswitch create function
type VswitchGrantCreateBody struct {
	Userid string `json:"name"`
	Nic string `json:"rdev"`
	Vswitch string `json:"vswitch"`
}

func getEndpointForCouple(endpoint string, userid string, nic string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/guests/")
	buffer.WriteString(userid)
	buffer.WriteString("/nic/")
	buffer.WriteString(nic)

        return buffer
}

func getEndpointForGrant(endpoint string, vswitch string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/vswitches/")
        buffer.WriteString(vswitch)

        return buffer
}

func buildVswitchGrantRequest(body *grantBody) ([]byte) {
	data, _ := json.Marshal(*body)

        return data
}

func buildVswitchCoupleRequest(cb *coupleBody) ([]byte) {
	data, _ := json.Marshal(*cb)

        return data
}

func vswitchCoupleCreate(endpoint string, body VswitchGrantCreateBody, ctxt RequestContext) (int, []byte) {
	var cb coupleBody
	cb.Info.Vswitch = body.Vswitch
	cb.Info.Couple = 1
	
	b := buildVswitchCoupleRequest(&cb)

	buffer := getEndpointForCouple(endpoint, body.Userid, body.Nic)
	status, data := put(buffer.String(), b, ctxt)

	return status, data
}

func vswitchGrantCreate(endpoint string, body VswitchGrantCreateBody, ctxt RequestContext) (int, []byte) {
	var gb grantBody
        gb.Vswitch.Userid = body.Userid

        b := buildVswitchGrantRequest(&gb)

        buffer := getEndpointForGrant(endpoint, body.Vswitch)
        status, data := put(buffer.String(), b, ctxt)

        return status, data
}
	
func VswitchGrant(endpoint string, name string, body VswitchGrantCreateBody) (int, []byte) {
	headers := buildAuthContext("")
	ctxt := RequestContext{
		values: headers,
	}

	// FIXME
	vswitchCoupleCreate(endpoint, body, ctxt)
	status, data := vswitchGrantCreate(endpoint, body, ctxt)
	return status, data
}
