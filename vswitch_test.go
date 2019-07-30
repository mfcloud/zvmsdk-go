package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_VswitchList(t *testing.T) {
	buf := getEndpointwithVswitchs(testEndpoint)
	hmock.On("Get", buf.String()).Return(200, []byte(""))

	status, _ := VswitchList(testEndpoint)
	require.Equal(t, 200, status)
}

func Test_VswitchDelete(t *testing.T) {
	buf := getEndpointwithVswitchs(testEndpoint)
	buf.WriteString("/id1")
	hmock.On("Delete", buf.String(), []byte(nil)).Return(200, []byte(""))
	status, _ := VswitchDelete(testEndpoint, "id1")
	require.Equal(t, 200, status)
}

func Test_VswitchGet(t *testing.T) {
	buf := getEndpointwithVswitchs(testEndpoint)
	buf.WriteString("/vsw1")
	hmock.On("Get", buf.String()).Return(200, []byte(""))

	status, _ := VswitchGet(testEndpoint, "vsw1")
	require.Equal(t, 200, status)
}

func Test_VswitchCreate(t *testing.T) {
        var vs VswitchCreateBody 
        vs.Name = "name1"
        vs.Rdev = "1"
        buf := getEndpointwithVswitchs(testEndpoint)
        body := `{"name":"name1","rdev":"1"}`
        hmock.On("Post", buf.String(), []byte(body)).Return(200, []byte(""))

        status, _ := VswitchCreate(testEndpoint, vs)
        require.Equal(t, 200, status)
}
