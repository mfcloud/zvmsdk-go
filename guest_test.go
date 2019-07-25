package zvmsdk

import (
	"testing"
	"github.com/stretchr/testify/require"

)

var hmock *HttpRequestMock

func init() {
	hmock = new(HttpRequestMock)
	hq = hmock
}

func Test_GuestList(t *testing.T) {
        //FIXME: mock this later
        status, _ := GuestList(testEndpoint)
        require.Equal(t, 200, status)
}

func Test_GuestGet(t *testing.T) {
	name := "name1"

        status, _ := GuestGet(testEndpoint, name)
        require.Equal(t, 200, status)
}

func Test_GuestGetInfo(t *testing.T) {
        name := "name1"

        status, _ := GuestGetInfo(testEndpoint, name)
        require.Equal(t, 200, status)
}

func Test_GuestGetPowerState(t *testing.T) {
        name := "name1"

        status, _ := GuestGetPowerState(testEndpoint, name)
        require.Equal(t, 200, status)
}

func Test_GuestCreate(t *testing.T) {
        disklist := make(GuestCreateDiskList, 2)
        var vs GuestCreateBody
        vs.Userid = "name1"
        vs.Vcpus = 1
        vs.Memory = 32
        vs.DiskPool = "disk1"

        disklist[0].Size = "1G"
        disklist[0].Format = "ECKD"
        disklist[0].Boot = "1"
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = "0"

        vs.DiskList = disklist

	hmock.On("Post", "1234").Return(nil)

        status, _ := GuestCreate(testEndpoint, vs)
	require.Equal(t, 200, status)
}

func Test_GuestCreateDisk(t *testing.T) {
	disklist := make(GuestCreateDiskList, 2)

        disklist[0].Size = "1G"
        disklist[0].Format = "ECKD"
        disklist[0].Boot = "1"
        disklist[1].Size = "2G"
        disklist[1].Format = "FBA"
        disklist[1].Boot = "0"

        status, _ := GuestCreateDisks(testEndpoint, "name1", disklist)
        require.Equal(t, 200, status)
}


func Test_GuestDeleteDisk(t *testing.T) {
        body := GuestDeleteDiskBody{}
	body.VdevList = make([]string, 2)
	body.VdevList[0] = "123"
	body.VdevList[1] = "456"

        status, _ := GuestDeleteDisks(testEndpoint, "name1", body)
	require.Equal(t, 200, status)
}

func Test_GuestConfigDisk(t *testing.T) {
        disklist := make(GuestConfigDiskList, 2)

        disklist[0].Vdev = "1111"
        disklist[0].Format = "ECKD"
        disklist[0].MntDir = "/mnt1"
        disklist[1].Vdev = "2222"
        disklist[1].Format = "FBA"
        disklist[1].MntDir = "/mnt/abc"

        status, _ := GuestConfigDisks(testEndpoint, "name1", disklist)
        require.Equal(t, 200, status)
}
