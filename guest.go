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
	Boot string `json:"is_boot_disk,omitempty"`
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
	Userid string `json:"userid"`
	Vcpus int `json:"vcpus,omitempty"`
	Memory int `json:"memory,omitempty"`
	DiskList GuestCreateDiskList `json:"disk_list,omitempty"`
	DiskPool string `json:"disk_pool,omitempty"`
	UserProfile string `json:"user_profile,omitempty"`
}

type GuestCreateBodyWrapper struct {
	Guest GuestCreateBody `json:"guest,omitempty"`
}

// GuestDeployBody will be used by upper layer
// when calling guest create function
type GuestDeployBody struct {
	Action string `json:"action"`
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

func buildGuestCreateRequest(body GuestCreateBodyWrapper) ([]byte) {
	data, _ := json.Marshal(body)
	return data
}

func buildGuestConfigDiskRequest(disklist GuestConfigDiskList) ([]byte) {
        data, _ := json.Marshal(disklist)

        return data
}

func buildGuestCreateNicRequest(body GuestCreateNicBody) ([]byte) {
	data, _ := json.Marshal(body)

        return data
}

// GuestCreate creates a guest
func GuestCreate(endpoint string, body GuestCreateBody) (int, []byte) {
	var gc GuestCreateBodyWrapper
	gc.Guest = body
	b := buildGuestCreateRequest(gc)

	buffer := getEndpointwithGuests(endpoint)
	status, data := hq.Post(buffer.String(), b)

	return status, data
}

// GuestList lists the guests on the host (z/VM)
func GuestList(endpoint string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
        status, data := hq.Get(buffer.String())

        return status, data
}

// GuestDelete deletes a guest
func GuestDelete(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

        status, data := hq.Delete(buffer.String(), nil)

	return status, data
}

func buildGuestDeployRequest(body GuestDeployBody) ([]byte) {
	data, _ := json.Marshal(body)
        return data
}

// GuestDeploy deploy an image to a given guest
func GuestDeploy(endpoint string, userid string, body GuestDeployBody) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(userid)
	buffer.WriteString("/")
	buffer.WriteString("action")

	body.Action = "deploy"
	b := buildGuestDeployRequest(body)
        status, data := hq.Post(buffer.String(), b)

        return status, data
}

// GuestGet retrieves user directory definition from a guest
func GuestGet(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
	buffer.WriteString("/")
        buffer.WriteString(guestid)

        status, data := hq.Get(buffer.String())

        return status, data
}

// GuestGetInfo gets information from guest
func GuestGetInfo(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
	buffer.WriteString("/info")

        status, data := hq.Get(buffer.String())

        return status, data
}

// GuestGetNic gets NIC information
func GuestGetNic(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/nic")

        status, data := hq.Get(buffer.String())

        return status, data
}

// GuestCreateDisks creates disk(s) on a given guest
func GuestCreateDisks(endpoint string, guestid string, body GuestCreateDiskList) (int, []byte) {

	createReq, _ := json.Marshal(body)

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        status, data := hq.Post(buffer.String(), createReq)

        return status, data
}

// GuestDeleteDisks deletes disk(s) from a guest
func GuestDeleteDisks(endpoint string, guestid string, body GuestDeleteDiskBody) (int, []byte) {
        deleteReq, _ := json.Marshal(body)

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        status, data := hq.Delete(buffer.String(), deleteReq)

        return status, data
}

// GuestConfigDisks configure disks on a guest
func GuestConfigDisks(endpoint string, guestid string, body GuestConfigDiskList) (int, []byte) {
	putReq := buildGuestConfigDiskRequest(body)

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/disks")

        headers := buildAuthContext("")
        ctxt := RequestContext{
                values: headers,
        }

        status, data := hq.Put(buffer.String(), putReq, ctxt)

        return status, data
}

// GuestsGetNics gets NIC info from guest
func GuestsGetNics(endpoint string) (int, []byte) {
        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/nics")

        status, data := hq.Get(buffer.String())

        return status, data
}

// GuestsGetVnics get VNIC information from all guests
func GuestsGetVnics(endpoint string) (int, []byte) {
        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/vnicsinfo")

        status, data := hq.Get(buffer.String())

        return status, data
}

// GuestsGetStats get stats from all guests
func GuestsGetStats(endpoint string) (int, []byte) {
        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/stats")

        status, data := hq.Get(buffer.String())

        return status, data
}

func GuestGetPowerState(endpoint string, guestid string) (int, []byte) {

        buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString("/")
        buffer.WriteString(guestid)
        buffer.WriteString("/power_state")

        status, data := hq.Get(buffer.String())

        return status, data
}
