package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_VswitchCoupleGrantCreate(t *testing.T) {
	var vs VswitchGrantCreateBody
	vs.Userid = "name1"
	vs.Nic = "1000"
	vs.Vswitch = "vsw1"
	headers := buildAuthContext("")
	ctxt := RequestContext{
		values: headers,
	}

	buf := getEndpointwithGuests(testEndpoint)
	buf.WriteString("/name1/nic/1000")
	body := `{"info":{"couple":"1","active":"0","vswitch":"vsw1"}}`
	hmock.On("Put", buf.String(), []byte(body), ctxt).Return(200, []byte(""))

	buf = getEndpointForGrant(testEndpoint, "vsw1")
	body = `{"vswitch":{"grant_userid":"name1"}}`
	hmock.On("Put", buf.String(), []byte(body), ctxt).Return(200, []byte(""))

	status, _ := VswitchCoupleGrant(testEndpoint, vs)
	require.Equal(t, 200, status)
}
