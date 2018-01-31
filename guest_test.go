package zvmsdk

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/require"

)


func Test_buildGuestCreateRequestJSON(t *testing.T) {
	disklist := make(GuestCreateDiskStructList, 2)
	var vs GuestCreateBodyStruct
	vs.Userid = "name1"
	vs.Vcpus = 1
	vs.Memory = 32
	vs.DiskPool = "disk1"

	disklist[0].Size = "1G"
	disklist[0].Format = "ECKD"
	disklist[0].Boot = 1
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = 0

	vs.DiskList = disklist

	//user_profile is omitted to test optional

	data := buildGuestCreateRequestJSON(vs)

	result := GuestCreateBodyStruct{}
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
        status, _ := GuestList(testEndpoint)
        require.Equal(t, status, 200)
}

func Test_GuestGet(t *testing.T) {
	name := "name1"

        status, _ := GuestGet(testEndpoint, name)
        require.Equal(t, status, 200)
}

func Test_GuestGetInfo(t *testing.T) {
        name := "name1"

        status, _ := GuestGetInfo(testEndpoint, name)
        require.Equal(t, status, 200)
}

func Test_GuestGetPowerState(t *testing.T) {
        name := "name1"

        status, _ := GuestGetPowerState(testEndpoint, name)
        require.Equal(t, status, 200)}



func Test_GuestCreate(t *testing.T) {
        disklist := make(GuestCreateDiskStructList, 2)
        var vs GuestCreateBodyStruct
        vs.Userid = "name1"
        vs.Vcpus = 1
        vs.Memory = 32
        vs.DiskPool = "disk1"

        disklist[0].Size = "1G"
        disklist[0].Format = "ECKD"
        disklist[0].Boot = 1
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = 0

        vs.DiskList = disklist

        status, _ := GuestCreate(testEndpoint, vs)
	require.Equal(t, status, 200)
}

func Test_GuestCreateNic(t *testing.T) {
        var vs GuestCreateNicBody
        vs.Vdev = "1111"
        vs.NicID = "id1"
        vs.MacAddr = "addr1"
        vs.Active = true

        status, _ := GuestCreateNic(testEndpoint, "name1", vs)
        require.Equal(t, status, 200)
}

func Test_GuestGetNic(t *testing.T) {
        name := "name1"

        status, _ := GuestGetNic(testEndpoint, name)
        require.Equal(t, status, 200)
}



func Test_GuestCreateDisk(t *testing.T) {
	disklist := make(GuestCreateDiskStructList, 2)

        disklist[0].Size = "1G"
        disklist[0].Format = "ECKD"
        disklist[0].Boot = 1
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = 0

        status, _ := GuestCreateDisks(testEndpoint, "name1", disklist)
        require.Equal(t, status, 200)
}


func Test_GuestDeleteDisk(t *testing.T) {
        body := GuestDeleteDiskBodyStruct{}
	body.VdevList = make([]string, 2)
	body.VdevList[0] = "123"
	body.VdevList[1] = "456"

        status, _ := GuestDeleteDisks(testEndpoint, "name1", body)
	require.Equal(t, status, 200)
}

func Test_GuestConfigDisk(t *testing.T) {
        disklist := make(GuestConfigDiskStructList, 2)

        disklist[0].Vdev = "1111"
        disklist[0].Format = "ECKD"
        disklist[0].MntDir = "/mnt1"
        disklist[1].Vdev = "2222"
        disklist[1].Format = "FBA"
        disklist[1].MntDir = "/mnt/abc"

        status, _ := GuestConfigDisks(testEndpoint, "name1", disklist)
        require.Equal(t, status, 200)
}

func Test_GuestGetNics(t *testing.T) {
        status, _ := GuestsGetNics(testEndpoint)
        require.Equal(t, status, 200)
}

func Test_GuestGetVnics(t *testing.T) {
        status, _ := GuestsGetVnics(testEndpoint)
        require.Equal(t, status, 200)
}

func Test_GuestGetStats(t *testing.T) {
        status, _ := GuestsGetStats(testEndpoint)
        require.Equal(t, status, 200)
}

