package zvmsdk


import (
	"bytes"
	"encoding/json"
)


type GuestCreateDiskList struct {
	Size string `json:"size"`
	Format string `json:"format"`
	Boot int32 `json:"is_boot_disk"`
}

type GuestCreateBody struct {
	Userid string `json:"userid"`
	Vcpus int `json:"vcpus"`
	Memory int `json:"memory"`
	Disklist []GuestCreateDiskList `json:"disk_list"`
	Diskpool string `json:"disk_pool"`
	Userprofile string `json:"user_profile"`
}


func getEndpointwithGuests(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/guests")
        return buffer
}


func buildGuestCreateDiskListJson(disklist []GuestCreateDiskList) ([]map[string]interface{}) {
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
	mkeys := []string{"userid", "vcpus", "memory"}
        mvalues := []interface{}{body.Userid, body.Vcpus, body.Memory}

        okeys := []string{"disk_pool", "user_profile", "disk_list"}
	disklist := buildGuestCreateDiskListJson(body.Disklist)
        ovalues := []interface{}{body.Diskpool, body.Userprofile, disklist}

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
        buffer.WriteString(guestid)

        status, data := delete(buffer.String(), nil)

	return status, data
}

func buildJson(keys []string, values []interface{}) ([]byte) {
        // map values to keys
        m := make(map[string]interface{})

        for i,v := range values {
                m[keys[i]] = v
        }

        // convert map to JSON
        data, _ := json.Marshal(m)

        return data

}

func buildGuestDeployRequestJson(image string, vdev string) ([]byte) {
        keys := []string{"image", "vdev"}
        values := []interface{}{image, vdev}

	return buildJson(keys, values)
}


func GuestDeploy(endpoint string, guestid string, image string, vdev string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString(guestid)

	body := buildGuestDeployRequestJson(image, vdev)
        status, data := post(buffer.String(), body)

        return status, data
}

func GuestQuery(endpoint string, guestid string) (int, []byte) {

	buffer := getEndpointwithGuests(endpoint)
        buffer.WriteString(guestid)

        status, data := get(buffer.String())

        return status, data
}

