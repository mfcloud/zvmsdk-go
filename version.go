package zvmsdk

import (
	"bytes"
)

func getEndpointwithVersion(endpoint string) bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(endpoint)
	return buffer
}

// VersionGet is to get version of z/VM cloud connector
func VersionGet(endpoint string) (int, []byte) {

	buffer := getEndpointwithVersion(endpoint)
	status, data := hq.Get(buffer.String())

	return status, data
}
