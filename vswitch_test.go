package zvmsdk

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/require"

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

func Test_buildVswitchUpdateRequest(t *testing.T) {
        var vs VswitchUpdateBody
        vs.GrantUserId = "id1"

        data := buildVswitchUpdateRequest(vs)

        result := VswitchUpdateBody{}
        err := json.Unmarshal(data, &result)
        if err != nil {
                panic(err.Error())
        }
        require.Equal(t, result.GrantUserId, "id1")
}


func Test_VswitchList(t *testing.T) {
	status, _ := VswitchList(test_endpoint)
	require.Equal(t, status, 200)
}

func Test_VswitchDelete(t *testing.T) {
	status, _ := VswitchDelete(test_endpoint, "id1")
	require.Equal(t, status, 200)
}

func Test_VswitchUpdate(t *testing.T) {
	var vs VswitchUpdateBody

	vs.GrantUserId = "id1"
        status, _ := VswitchUpdate(test_endpoint, "vsw1", vs)
        require.Equal(t, status, 200)
}
