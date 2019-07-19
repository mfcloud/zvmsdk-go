package zvmsdk


import (
	"bytes"
	"encoding/json"
)


// GuestCreateDisk will be used by upper layer
// when calling guest create disk function
type GuestCreateDisk struct {
	Size string `json:"size,omitempty"`
	Format string `json:"format,omitempty"`
	Boot int32 `json:"is_boot_disk,omitempty"`
}

// GuestCreateDiskList will be used by upper layer
// when calling guest create disk function
type GuestCreateDiskList []GuestCreateDisk

// GuestConfigDisk will be used by upper layer when
// calling guest create disk function
type GuestConfigDisk struct {
        Vdev string `json:"vdev"`
        Format string `json:"format"`
        MntDir string `json:"mntdir"`
}

// GuestConfigDiskList will be used by upper layer
// when calling guest create disk function
type GuestConfigDiskList []GuestConfigDisk

// GuestCreateBody will be used by upper layer
// when calling guest create function
type GuestCreateBody struct {
	Userid string `json:"userid,omitempty"`
	Vcpus int `json:"vcpus,omitempty"`
	Memory int `json:"memory,omitempty"`
	DiskList GuestCreateDiskList `json:"disk_list,omitempty"`
	DiskPool string `json:"disk_pool,omitempty"`
	UserProfile string `json:"user_profile,omitempty"`
}

// GuestDeployBody will be used by upper layer
// when calling guest create function
type GuestDeployBody struct {
        Userid string `json:"userid,omitempty"`
	Image string `json:"image,omitempty"`
	TransportFiles string `json:"transportfiles,omitempty"`
	RemoteHost string `json:"remotehost,omitempty"`
	Vdev string `json:"vdev,omitempty"`
}

// GuestDeleteVdevList will be used by upper layer
// when calling guest delete disk function
type GuestDeleteVdevList []string

// GuestDeleteDiskBody will be used by upper layer
// calling guest delete disk function
type GuestDeleteDiskBody struct {
	VdevList GuestDeleteVdevList `json:"vdev_list"`
}

// GuestCreateNicBody will be used by upper layer when calling
// guest create nic function
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

func buildGuestCreateDiskListJSON(disklist GuestCreateDiskList) ([]map[string]interface{}) {
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


func buildGuestCreateRequestJSON(body GuestCreateBody) ([]byte) {
	mkeys := []string{"userid", "vcpus", "memory", "user_profile", "disk_pool"}
        mvalues := []interface{}{body.Userid, body.Vcpus, body.Memory, body.UserProfile, body.DiskPool}

        okeys := []string{"disk_list"}
	disklist := buildGuestCreateDiskListJSON(body.DiskList)
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

        return buildJSON(keys, values)
}

func buildGuestConfigDiskRequest(disklist GuestConfigDiskList) ([]byte) {
        length := len(disklist)

        ret := make([]map[string]interface{}, length)

        keys := []string{"vdev", "format", "mntdir"}

        for k, element := range disklist {
                ret[k] = make(map[string]interface{})
                ovalues := []interface{}{element.Vdev, element.Format, element.MntDir}
                for i,v := range ovalues {
                        ret[k][keys[i]] = v
                }
        }
        data, _ := json.Marshal(ret)

        return data
}

func buildGuestCreateNicRequestJSON(body GuestCreateNicBody) ([]byte) {
        keys := []string{"vdev", "nic_id", "mac_addr", "active"}
        values := []interface{}{body.Vdev, body.NicID, body.MacAddr, body.Active}

        return buildJSON(keys, values)
}

// GuestCreate creates a guest
func GuestCreate(endpoint string, body GuestCreateBody) (int, []byte) {

	createJSON := buildGuestCreateRequestJSON(body)

	buffer := getEndpointwithGuests(endpoint)
	status, data := post(buffer.String(), createJSON)

	return status, data
}

// GuestList lists the guests on the host (z/VM)
func GuestList(endpoint string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
        status, data := get(buffer.String())

        return status, data
}

// GuestDelete deletes a guest
func GuestDelete(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

        status, data := delete(buffer.String(), nil)

	return status, data
}

func buildGuestDeployRequestJSON(action string, body GuestDeployBody) ([]byte) {
        keys := []string{"action", "image", "vdev", "transportfiles", "remotehost"}
        values := []interface{}{action, body.Image, body.Vdev, body.TransportFiles, body.RemoteHost}

	return buildJSON(keys, values)
}

// GuestDeploy deploy an image to a given guest
func GuestDeploy(endpoint string, body GuestDeployBody) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(body.Userid)

	b := buildGuestDeployRequestJSON("deploy", body)
        status, data := post(buffer.String(), b)

        return status, data
}

// GuestGet retrieves user directory definition from a guest
func GuestGet(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

        status, data := get(buffer.String())

        return status, data
}

// GuestGetInfo gets information from guest
func GuestGetInfo(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
	buffer.WriteString("/info")

        status, data := get(buffer.String())

        return status, data
}

// GuestGetNic gets NIC information
func GuestGetNic(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/nic")

        status, data := get(buffer.String())

        return status, data
}

// GuestCreateNic create NIC
func GuestCreateNic(endpoint string, guestid string, body GuestCreateNicBody) (int, []byte) {

        createJSON := buildGuestCreateNicRequestJSON(body)

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/nic")

        status, data := post(buffer.String(), createJSON)

        return status, data
}

// GuestGetPowerState gets power state of a guest
func GuestGetPowerState(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/power_state")

        status, data := get(buffer.String())

        return status, data
}

// GuestCreateDisks creates disk(s) on a given guest
func GuestCreateDisks(endpoint string, guestid string, body GuestCreateDiskList) (int, []byte) {

	createReq, _ := json.Marshal(buildGuestCreateDiskListJSON(body))

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        status, data := post(buffer.String(), createReq)

        return status, data
}

// GuestDeleteDisks deletes disk(s) from a guest
func GuestDeleteDisks(endpoint string, guestid string, body GuestDeleteDiskBody) (int, []byte) {
        deleteReq, _ := json.Marshal(buildGuestDeleteDiskRequest(body))

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        status, data := delete(buffer.String(), deleteReq)

        return status, data
}

// GuestConfigDisks configure disks on a guest
func GuestConfigDisks(endpoint string, guestid string, body GuestConfigDiskList) (int, []byte) {
	putReq := buildGuestConfigDiskRequest(body)

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        headers := buildAuthContext("abc")
        ctxt := RequestContext{
                values: headers,
        }

        status, data := put(buffer.String(), putReq, ctxt)

        return status, data
}

// GuestsGetNics gets NIC info from guest
func GuestsGetNics(endpoint string) (int, []byte) {
        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/nics")

        status, data := get(buffer.String())

        return status, data
}

// GuestsGetVnics get VNIC information from all guests
func GuestsGetVnics(endpoint string) (int, []byte) {
        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/vnicsinfo")

        status, data := get(buffer.String())

        return status, data
}

// GuestsGetStats get stats from all guests
func GuestsGetStats(endpoint string) (int, []byte) {
        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/stats")

        status, data := get(buffer.String())

        return status, data
}
