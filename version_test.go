package zvmsdk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_VersionGet(t *testing.T) {

	status, _ := VersionGet(testEndpoint)

	require.Equal(t, status, 200)
}
