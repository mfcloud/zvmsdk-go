package zvmsdk


// VersionGet is to get version of z/VM cloud connector
func VersionGet(endpoint string) (int, []byte) {

	status, data := hq.Get(endpoint)

	return status, data
}

