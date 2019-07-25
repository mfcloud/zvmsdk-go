package zvmsdk


import (
	"bytes"
	"encoding/json"
)


type coupleInfo struct {
	Couple string `json:"couple"`
	Active string `json:"active"`
	Vswitch string `json:"vswitch"`
}

type coupleBody struct {
	Info   coupleInfo `json:"info,omitempty"`
}

type grantUserid struct {
	Userid string `json:"grant_userid,omitempty"`
}

type grantBody struct {
	Vswitch grantUserid `json:"vswitch,omitempty"`
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

func vswitchCouple(endpoint string, body VswitchGrantCreateBody, ctxt RequestContext) (int, []byte) {
	var cb coupleBody
	cb.Info.Vswitch = body.Vswitch
	cb.Info.Couple = "1"
	cb.Info.Active = "0"

	b := buildVswitchCoupleRequest(&cb)

	buffer := getEndpointForCouple(endpoint, body.Userid, body.Nic)
	status, data := hq.Put(buffer.String(), b, ctxt)

	return status, data
}

func vswitchGrant(endpoint string, body VswitchGrantCreateBody, ctxt RequestContext) (int, []byte) {
	var gb grantBody
        gb.Vswitch.Userid = body.Userid

        b := buildVswitchGrantRequest(&gb)

        buffer := getEndpointForGrant(endpoint, body.Vswitch)
        status, data := hq.Put(buffer.String(), b, ctxt)

        return status, data
}
	
func VswitchGrant(endpoint string, body VswitchGrantCreateBody) (int, []byte) {
	headers := buildAuthContext("")
	ctxt := RequestContext{
		values: headers,
	}

	// FIXME
	vswitchCouple(endpoint, body, ctxt)
	status, data := vswitchGrant(endpoint, body, ctxt)
	return status, data
}
