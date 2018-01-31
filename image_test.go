package zvmsdk

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/require"
)


func Test_buildImageCreateRequest(t *testing.T) {

	meta := map[string]string{"apple": "5", "lettuce": "7"}
	data := buildImageCreateRequest("name", "url", meta, "host")

	s := ImageCreateBody{}
	err := json.Unmarshal(data, &s)
        if err != nil {
		panic(err.Error())
	}
	require.Equal(t, s.Name, "name")
	require.Equal(t, s.URL, "url")
	require.Equal(t, s.RemoteHost, "host")
	require.Equal(t, s.Meta["apple"], "5")
	require.Equal(t, s.Meta["lettuce"], "7")
}

func Test_ImageGetRootDiskSize(t *testing.T) {
        status, _ := ImageGetRootDiskSize(testEndpoint, "image1")
        require.Equal(t, status, 200)
}

func Test_ImageGetWithName(t *testing.T) {
        status, _ := ImageGet(testEndpoint, "image1")
        require.Equal(t, status, 200)
}

func Test_ImageGet(t *testing.T) {
        status, _ := ImageGet(testEndpoint, "")
        require.Equal(t, status, 200)
}

func Test_buildImageUpdateRequest(t *testing.T) {
        data := buildImageUpdateRequest("url1", "host1")

        s := ImageUpdateBody{}
        err := json.Unmarshal(data, &s)
        if err != nil {
                panic(err.Error())
        }
        require.Equal(t, s.DestURL, "url1")
        require.Equal(t, s.RemoteHost, "host1")
}



func Test_ImageCreate(t *testing.T) {
        var ic ImageCreateBody
	meta := map[string]string{"apple": "5", "lettuce": "7"}

        ic.Name = "name1"
        ic.RemoteHost = "remotehost1"
	ic.Meta = meta
	ic.URL= "url1"

        status, _ := ImageCreate(testEndpoint, ic)
        require.Equal(t, status, 200)
}

func Test_ImageUpdate(t *testing.T) {
        var ic ImageUpdateBody

        ic.DestURL= "url1"
	ic.RemoteHost = "name1"

        status, _ := ImageUpdate(testEndpoint, "vsw1", ic)
        require.Equal(t, status, 200)

}