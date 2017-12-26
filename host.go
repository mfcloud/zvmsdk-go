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



func HostInfo(endpoint string) (int, []byte) {
	buffer := getEndpointwithHost(endpoint)
	status, data := get(buffer.String())

	return status, data
}

func HostDiskpoolInfo(endpoint string, disk string) (int, []byte) {

	buffer := getEndpointwithHost(endpoint)
	buffer.WriteString("/disk/")
        buffer.WriteString(disk)

        status, data := get(buffer.String())

        return status, data
}
