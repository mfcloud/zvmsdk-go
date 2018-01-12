package zvmsdk


import (
	"bytes"
	"encoding/json"
)


type GuestCreateDiskStruct struct {
	Size string `json:"size"`
	Format string `json:"format"`
	Boot int32 `json:"is_boot_disk"`
}

type GuestCreateDiskStructList []GuestCreateDiskStruct

type GuestCreateBody struct {
	Userid string `json:"userid"`
	Vcpus int `json:"vcpus"`
	Memory int `json:"memory"`
	DiskList GuestCreateDiskStructList `json:"disk_list"`
	DiskPool string `json:"disk_pool"`
	UserProfile string `json:"user_profile"`
}

type GuestDeleteVdevList []string

type GuestDeleteDiskBody struct {
	VdevList GuestDeleteVdevList `json:"vdev_list"`
}

type GuestCreateNicBody struct {
        Vdev string `json:"vdev"`
        NicID string `json:"nic_id"`
        MacAddr string `json:"mac_addr"`
	Active bool `json:"active"`
}

func getEndpointwithGuests(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/guests")
        return buffer
}


func buildGuestCreateDiskListJson(disklist GuestCreateDiskStructList) ([]map[string]interface{}) {
	length := len(disklist)

	ret := make([]map[string]interface{}, length)

	mkeys := []string{"size", "format", "is_boot_disk"}

	for k, element := range disklist {
		ret[k] = make(map[string]interface{})
	        ovalues := []interface{}{element.Size, element.Format, element.Boot}
	        for i,v := range ovalues {
	                ret[k][mkeys[i]] = v
	        }
	}
	return ret
}


func buildGuestCreateRequestJson(body GuestCreateBody) ([]byte) {
	mkeys := []string{"userid", "vcpus", "memory", "user_profile", "disk_pool"}
        mvalues := []interface{}{body.Userid, body.Vcpus, body.Memory, body.UserProfile, body.DiskPool}

        okeys := []string{"disk_list"}
	disklist := buildGuestCreateDiskListJson(body.DiskList)
        ovalues := []interface{}{disklist}

        // map values to keys
        m := make(map[string]interface{})

        for i,v := range mvalues {
                m[mkeys[i]] = v
        }

	// only non-empty value will be put into map
	for i, v := range ovalues {
		m[okeys[i]] = v
	}
        // convert map to JSON
        data, _ := json.Marshal(m)

	return data
}

func buildGuestDeleteDiskRequest(body GuestDeleteDiskBody) ([]byte) {
        keys := []string{"vdev_list"}
        values := []interface{}{body.VdevList}

        return buildJson(keys, values)
}


func buildGuestCreateNicRequestJson(body GuestCreateNicBody) ([]byte) {
        keys := []string{"vdev", "nic_id", "mac_addr", "active"}
        values := []interface{}{body.Vdev, body.NicID, body.MacAddr, body.Active}

        return buildJson(keys, values)
}


func GuestCreate(endpoint string, body GuestCreateBody) (int, []byte) {

	createJson := buildGuestCreateRequestJson(body)

	buffer := getEndpointwithGuests(endpoint)
	status, data := post(buffer.String(), createJson)

	return status, data
}

func GuestList(endpoint string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
        status, data := get(buffer.String())

        return status, data
}


func GuestDelete(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

        status, data := delete(buffer.String(), nil)

	return status, data
}

func buildGuestDeployRequestJson(image string, vdev string) ([]byte) {
        keys := []string{"image", "vdev"}
        values := []interface{}{image, vdev}

	return buildJson(keys, values)
}


func GuestDeploy(endpoint string, guestid string, image string, vdev string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

	body := buildGuestDeployRequestJson(image, vdev)
        status, data := post(buffer.String(), body)

        return status, data
}

func GuestGet(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

        status, data := get(buffer.String())

        return status, data
}

func GuestGetInfo(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
	buffer.WriteString("/info")

        status, data := get(buffer.String())

        return status, data
}

func GuestGetNic(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/nic")

        status, data := get(buffer.String())

        return status, data
}

func GuestCreateNic(endpoint string, guestid string, body GuestCreateNicBody) (int, []byte) {

        createJson := buildGuestCreateNicRequestJson(body)

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/nic")

        status, data := post(buffer.String(), createJson)

        return status, data
}


func GuestGetPowerState(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/power_state")

        status, data := get(buffer.String())

        return status, data
}

func GuestCreateDisks(endpoint string, guestid string, body GuestCreateDiskStructList) (int, []byte) {

	createReq, _ := json.Marshal(buildGuestCreateDiskListJson(body))

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        status, data := post(buffer.String(), createReq)

        return status, data
}

func GuestDeleteDisks(endpoint string, guestid string, body GuestDeleteDiskBody) (int, []byte) {
        deleteReq, _ := json.Marshal(buildGuestDeleteDiskRequest(body))

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        status, data := delete(buffer.String(), deleteReq)

        return status, data
}

