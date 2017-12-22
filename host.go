package zvmsdk

import (
	"bytes"
)


func HostInfo() (int, []byte) {
	status, data := get("http://localhost:8080/host")

	return status, data
}

func HostDiskpoolInfo(disk string) (int, []byte) {
        var buffer bytes.Buffer

        buffer.WriteString("http://localhost:8080/host/disk/")
        buffer.WriteString(disk)

        status, data := get(buffer.String())

        return status, data
}
