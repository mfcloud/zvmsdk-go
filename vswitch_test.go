package zvmsdk

import (
	"fmt"
	"testing"
	"encoding/json"

)


func Test_VswitchCreate(t *testing.T) {
	var vs VswitchCreateBody
	vs.Name = "name1"
	vs.Rdev = "rdev1"

	data := buildVswitchCreateRequest(vs)

	result := VswitchCreateBody{}
	err := json.Unmarshal(data, &result)
        if err != nil {
		panic(err.Error())
	}
	assertEqual(t, result.Name, "name1", "")
	assertEqual(t, result.Rdev, "rdev1", "")
	fmt.Println("Test_VswitchCreate passed")
}

func Test_VswitchUpdate(t *testing.T) {
        var vs VswitchUpdateBody
        vs.GrantUserId = "id1"

        data := buildVswitchUpdateRequest(vs)

        result := VswitchUpdateBody{}
        err := json.Unmarshal(data, &result)
        if err != nil {
                panic(err.Error())
        }
        assertEqual(t, result.GrantUserId, "id1", "")
	fmt.Println("Test_VswitchUpdate passed")
}

