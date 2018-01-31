package zvmsdk

import (
	"testing"
	"github.com/stretchr/testify/require"
)


func Test_TokenCreate(t *testing.T) {
	var v TokenCreateBody

	status, _ := TokenCreate(testEndpoint, v)

	require.Equal(t, status, 200)
}

