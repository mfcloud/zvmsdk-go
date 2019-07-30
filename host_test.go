package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_HostInfo(t *testing.T) {
	buf := getEndpointwithHost(testEndpoint)
	hmock.On("Get", buf.String()).Return(200, []byte(""))
	status, _ := HostInfo(testEndpoint)
	require.Equal(t, status, 200)
}

func Test_HostDiskPoolInfo(t *testing.T) {
	buf := getEndpointwithHost(testEndpoint)
	buf.WriteString("/disk/disk1")
	hmock.On("Get", buf.String()).Return(200, []byte(""))

	status, _ := HostDiskpoolInfo(testEndpoint, "disk1")
	require.Equal(t, status, 200)
}
