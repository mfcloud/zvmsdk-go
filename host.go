package zvmsdk

import (
	"bytes"
)


func getEndpointwithHost(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/host")
        return buffer
}


// HostInfo gets information for the host (z/VM) running on
func HostInfo(endpoint string) (int, []byte) {
	buffer := getEndpointwithHost(endpoint)
	status, data := get(buffer.String())

	return status, data
}

// HostDiskpoolInfo gets information for the disk pool exists on
// the host (z/VM) running on
func HostDiskpoolInfo(endpoint string, disk string) (int, []byte) {

	buffer := getEndpointwithHost(endpoint)
	buffer.WriteString("/disk/")
        buffer.WriteString(disk)

        status, data := get(buffer.String())

        return status, data
}
