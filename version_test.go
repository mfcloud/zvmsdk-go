package zvmsdk

import (
	"testing"
	"github.com/stretchr/testify/require"
)


func Test_VersionGet(t *testing.T) {

	status, _ := VersionGet(testEndpoint)

	require.Equal(t, status, 200)
}

