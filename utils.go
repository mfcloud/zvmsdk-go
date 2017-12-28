package zvmsdk


import (
        "encoding/json"
)


func buildJson(keys []string, values []interface{}) ([]byte) {

        // map values to keys
        m := make(map[string]interface{})
        for i,v := range values {
                m[keys[i]] = v
        }
        // convert map to JSON
        data, _ := json.Marshal(m)

        return data
}


func buildAuthContext(authToken string) (map[string]string) {
	m := make(map[string]string)

	m["X-Auth-Token"] = authToken
	return m
}

func buildAdminContext(adminToken string) (map[string]string) {
        m := make(map[string]string)

        m["X-Admin-Token"] = adminToken
        return m
}

