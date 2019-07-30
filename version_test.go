package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_VersionGet(t *testing.T) {
	buf := getEndpointwithVersion(testEndpoint)
	hmock.On("Get", buf.String()).Return(200, []byte(""))

	status, _ := VersionGet(testEndpoint)

	require.Equal(t, status, 200)
}
