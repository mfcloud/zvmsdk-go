package zvmsdk

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_buildVswitchCreateRequest(t *testing.T) {
	var vs VswitchCreateBody
	vs.Name = "name1"
	vs.Rdev = "rdev1"

	data := buildVswitchCreateRequest(vs)

	result := VswitchCreateBody{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		panic(err.Error())
	}
	require.Equal(t, result.Name, "name1")
	require.Equal(t, result.Rdev, "rdev1")
}

func Test_VswitchList(t *testing.T) {
	status, _ := VswitchList(testEndpoint)
	require.Equal(t, 200, status)
}

func Test_VswitchDelete(t *testing.T) {
	status, _ := VswitchDelete(testEndpoint, "id1")
	require.Equal(t, 200, status)
}

func Test_VswitchGet(t *testing.T) {
	status, _ := VswitchGet(testEndpoint, "vsw1")
	require.Equal(t, 200, status)
}
