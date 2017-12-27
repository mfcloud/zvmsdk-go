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
