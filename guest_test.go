package zvmsdk

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
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
	buf := getEndpointwithGuests(testEndpoint)
	body := `{"guest":{"userid":"name1","vcpus":1,"memory":32,"disk_list":[{"size":"1G","format":"ECKD","is_boot_disk":"1"},{"size":"2G","format":"FBA","is_boot_disk":"0"}],"disk_pool":"disk1"}}`
	hmock.On("Post", buf.String(), []byte(body)).Return(200, []byte(""))
	assert := assert.New(t)

	status, _ := GuestCreate(testEndpoint, vs)
	assert.Equal(200, status)
}

func Test_GuestCreateDisk(t *testing.T) {
	disklist := make(GuestCreateDiskList, 2)

	disklist[0].Size = "1G"
	disklist[0].Format = "ECKD"
	disklist[0].Boot = "1"
	disklist[1].Size = "2G"
	disklist[1].Format = "FBA"
	disklist[1].Boot = "0"

	buf := getEndpointwithGuests(testEndpoint)
	buf.WriteString("/name1/disks")
	body := `[{"size":"1G","format":"ECKD","is_boot_disk":"1"},{"size":"2G","format":"FBA","is_boot_disk":"0"}]`
	hmock.On("Post", buf.String(), []byte(body)).Return(200, []byte(""))
	status, _ := GuestCreateDisks(testEndpoint, "name1", disklist)
	require.Equal(t, 200, status)
}

func Test_GuestDeleteDisk(t *testing.T) {
	body := GuestDeleteDiskBody{}
	body.VdevList = make([]string, 2)
	body.VdevList[0] = "123"
	body.VdevList[1] = "456"

	buf := getEndpointwithGuests(testEndpoint)
	buf.WriteString("/name1/disks")
	b := `{"vdev_list":["123","456"]}`
	hmock.On("Delete", buf.String(), []byte(b)).Return(200, []byte(""))

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

	headers := buildAuthContext("")
	ctxt := RequestContext{
		values: headers,
	}

	buf := getEndpointwithGuests(testEndpoint)
	buf.WriteString("/name1/disks")
	b := `[{"vdev":"1111","format":"ECKD","mntdir":"/mnt1"},{"vdev":"2222","format":"FBA","mntdir":"/mnt/abc"}]`
	hmock.On("Put", buf.String(), []byte(b), ctxt).Return(200, []byte(""))
	status, _ := GuestConfigDisks(testEndpoint, "name1", disklist)
	require.Equal(t, 200, status)
}
