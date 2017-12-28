package zvmsdk

import (
	"bytes"
)


type TokenCreateBody struct {
	admin_token string
}

func getEndpointwithToken(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/token")
        return buffer
}

func TokenCreate(endpoint string, body TokenCreateBody) (int, []byte) {
	buffer := getEndpointwithToken(endpoint)

	status, data := post(buffer.String(), nil)

	return status, data
}

