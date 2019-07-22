package zvmsdk


import (
	"bytes"
	"encoding/json"
)


type GuestNetwork struct {
	IP string `json:"ip_addr,omitempty"`
	Cidr string `json:"cidr,omitempty"`
	Gateway string `json:"gateway_addr,omitempty"`
	Vdev string `json:"nic_vdev,omitempty"`
	Mac string `json:"mac_addr,omitempty"`
	NicID string `json:"nic_id,omitempty"`
	Osa string `json:"osa_device,omitempty"`
	Dns string `json:"dns_addr,omitempty"`
}

type GuestNetworkList []Network

type GuestInterface struct {
	Osversion string `json:"userid,omitempty"`
	Networks GuestNetworkList `json:"guest_networks,omitempty"`
}

// ImageCreateBody used as image create input param
type GuestInterfaceCreateBody struct {
	Userid string `json:"userid"`
	If Interface `json:"interface"`
}

func getEndpointwithInterface(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/interface")
        return buffer
}

func buildInterfaceCreateRequest(body GuestInterfaceCreateBody) ([]byte) {
	data, _ := json.Marshal(body)

        return data
}

func GuestInterfaceCreate(endpoint string, body GuestInterfaceCreateBody) (int, []byte) {

	request := buildnterfaceCreateRequest(body)

	buffer := getEndpointwithInterface(endpoint)
	status, data := post(buffer.String(), request)

	return status, data
}