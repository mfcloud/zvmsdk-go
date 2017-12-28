package zvmsdk


func VersionGet(endpoint string) (int, []byte) {

	status, data := get(endpoint)

	return status, data
}

