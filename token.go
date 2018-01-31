package zvmsdk

import (
	"bytes"
)


// TokenCreateBody is used for generating a token
type TokenCreateBody struct {
	adminToken string
}

func getEndpointwithToken(endpoint string) (bytes.Buffer) {
        var buffer bytes.Buffer

        buffer.WriteString(endpoint)
        buffer.WriteString("/token")
        return buffer
}

// TokenCreate creates a token with admin token
func TokenCreate(endpoint string, body TokenCreateBody) (int, []byte) {
	buffer := getEndpointwithToken(endpoint)

	status, data := post(buffer.String(), nil)

	return status, data
}

