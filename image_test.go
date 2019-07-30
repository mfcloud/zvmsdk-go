package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ImageGetRootDiskSize(t *testing.T) {
	buf := getEndpointwithImages(testEndpoint)
	buf.WriteString("/image1/root_disk_size")
	hmock.On("Get", buf.String()).Return(200, []byte(""))

	status, _ := ImageGetRootDiskSize(testEndpoint, "image1")
	require.Equal(t, status, 200)
}

func Test_ImageGetWithName(t *testing.T) {
	buf := getEndpointwithImages(testEndpoint)
	buf.WriteString("?imagename=image1")
	hmock.On("Get", buf.String()).Return(200, []byte(""))
	status, _ := ImageGet(testEndpoint, "image1")
	require.Equal(t, status, 200)
}

func Test_ImageGet(t *testing.T) {
	buf := getEndpointwithImages(testEndpoint)
	hmock.On("Get", buf.String()).Return(200, []byte(""))

	status, _ := ImageGet(testEndpoint, "")
	require.Equal(t, status, 200)
}

func Test_ImageCreate(t *testing.T) {
	var ic ImageCreateBody
	meta := map[string]string{"apple": "5", "lettuce": "7"}

	ic.Name = "name1"
	ic.RemoteHost = "remotehost1"
	ic.Meta = meta
	ic.URL = "url1"
	buf := getEndpointwithImages(testEndpoint)
	body := `{"image_name":"name1","remote_host":"remotehost1","image_meta":{"apple":"5","lettuce":"7"},"url":"url1"}`
	hmock.On("Post", buf.String(), []byte(body)).Return(200, []byte(""))

	status, _ := ImageCreate(testEndpoint, ic)
	require.Equal(t, status, 200)
}

func Test_ImageUpdate(t *testing.T) {
	var ic ImageUpdateBody

	ic.DestURL = "url1"
	ic.RemoteHost = "name1"

	headers := buildAuthContext("")
	ctxt := RequestContext{
		values: headers,
	}
	buf := getEndpointwithImages(testEndpoint)
	buf.WriteString("/image1")
	b := `{"dest_url":"url1","remote_host":"name1"}`
	hmock.On("Put", buf.String(), []byte(b), ctxt).Return(200, []byte(""))
	status, _ := ImageUpdate(testEndpoint, "image1", ic)
	require.Equal(t, status, 200)

}
