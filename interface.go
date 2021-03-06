package zvmsdk

import (
	"bytes"
	"encoding/json"
)

type GuestNetwork struct {
	IP      string `json:"ip_addr,omitempty"`
	Cidr    string `json:"cidr,omitempty"`
	Gateway string `json:"gateway_addr,omitempty"`
	Vdev    string `json:"nic_vdev,omitempty"`
	Mac     string `json:"mac_addr,omitempty"`
	NicID   string `json:"nic_id,omitempty"`
	Osa     string `json:"osa_device,omitempty"`
	Dns     string `json:"dns_addr,omitempty"`
}

type GuestNetworkList []GuestNetwork

type GuestInterface struct {
	Osversion string           `json:"os_version,omitempty"`
	Networks  GuestNetworkList `json:"guest_networks,omitempty"`
	Active    string           `json:"active,omityempty"`
}

// ImageCreateBody used as image create input param
type GuestInterfaceCreateBody struct {
	If GuestInterface `json:"interface"`
}

func getEndpointwithInterface(endpoint string, userid string) bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(endpoint)
	buffer.WriteString("/guests/")
	buffer.WriteString(userid)
	buffer.WriteString("/interface")
	return buffer
}

func buildInterfaceCreateRequest(body GuestInterfaceCreateBody) []byte {
	data, _ := json.Marshal(body)

	return data
}

func GuestInterfaceCreate(endpoint string, userid string, body GuestInterfaceCreateBody) (int, []byte) {

	request := buildInterfaceCreateRequest(body)

	buffer := getEndpointwithInterface(endpoint, userid)
	status, data := hq.Post(buffer.String(), request)

	return status, data
}
