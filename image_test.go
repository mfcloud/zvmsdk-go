package zvmsdk

import (
	"fmt"
	"testing"
	"encoding/json"

)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}


func Test_ImageCreate(t *testing.T) {
	t.Log("first pass")

	meta := map[string]string{"apple": "5", "lettuce": "7"}
	data := buildImageCreateRequest("name", "url", meta, "host")

	s := ImageCreateBody{}
	err := json.Unmarshal(data, &s)
        if err != nil {
		panic(err.Error())
	}
	assertEqual(t, s.Image_name, "name", "")
	assertEqual(t, s.Url, "url", "")
	assertEqual(t, s.Remote_host, "host", "")
	assertEqual(t, s.Image_meta["apple"], "5", "")
	assertEqual(t, s.Image_meta["lettuce"], "7", "")
}
