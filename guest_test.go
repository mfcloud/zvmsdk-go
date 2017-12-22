package zvmsdk

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/require"

)


func Test_buildGuestCreateRequestJson(t *testing.T) {
	disklist := make([]GuestCreateDiskList, 2)
	var vs GuestCreateBody
	vs.Userid = "name1"
	vs.Vcpus = 1
	vs.Memory = 32
	vs.Diskpool = "disk1"

	disklist[0].Size = "1G"
	disklist[0].Format = "ECKD"
	disklist[0].Boot = 1
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = 0

	vs.Disklist = disklist

	//user_profile is omitted to test optional

	data := buildGuestCreateRequestJson(vs)

	result := GuestCreateBody{}
	err := json.Unmarshal(data, &result)
        if err != nil {
		panic(err.Error())
	}
	require.Equal(t, result.Userid, "name1")
	require.Equal(t, result.Vcpus, 1)
	require.Equal(t, result.Memory, 32)
}

func Test_GuestList(t *testing.T) {
        //FIXME: mock this later
        status, _ := GuestList()
        require.Equal(t, status, 200)
}

func Test_GuestCreate(t *testing.T) {
        disklist := make([]GuestCreateDiskList, 2)
        var vs GuestCreateBody
        vs.Userid = "name1"
        vs.Vcpus = 1
        vs.Memory = 32
        vs.Diskpool = "disk1"

        disklist[0].Size = "1G"
        disklist[0].Format = "ECKD"
        disklist[0].Boot = 1
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = 0

        vs.Disklist = disklist

        status, _ := GuestCreate(vs)
	require.Equal(t, status, 200)
}

