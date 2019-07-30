package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_InterfaceCreate(t *testing.T) {
	var ic GuestInterfaceCreateBody

	ic.If.Osversion = "1"
	ic.If.Active = "0"
	ic.If.Networks = make(GuestNetworkList, 1)
	ic.If.Networks[0].IP = "1.2.3.4"
	ic.If.Networks[0].Vdev = "100"
	body := `{"interface":{"os_version":"1","guest_networks":[{"ip_addr":"1.2.3.4","nic_vdev":"100"}],"active":"0"}}`
	buf := getEndpointwithInterface(testEndpoint, "name1")
	hmock.On("Post", buf.String(), []byte(body)).Return(200, []byte(""))

	status, _ := GuestInterfaceCreate(testEndpoint, "name1", ic)
	require.Equal(t, status, 200)
}
