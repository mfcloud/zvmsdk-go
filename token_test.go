package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_TokenCreate(t *testing.T) {
	var v TokenCreateBody

	buf := getEndpointwithToken(testEndpoint)
	hmock.On("Post", buf.String(), []byte(nil)).Return(200, []byte(""))
	status, _ := TokenCreate(testEndpoint, v)

	require.Equal(t, status, 200)
}
