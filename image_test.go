package zvmsdk

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/require"
)


func Test_ImageCreate(t *testing.T) {

	meta := map[string]string{"apple": "5", "lettuce": "7"}
	data := buildImageCreateRequest("name", "url", meta, "host")

	s := ImageCreateBody{}
	err := json.Unmarshal(data, &s)
        if err != nil {
		panic(err.Error())
	}
	require.Equal(t, s.Name, "name")
	require.Equal(t, s.Url, "url")
	require.Equal(t, s.Remote_host, "host")
	require.Equal(t, s.Meta["apple"], "5")
	require.Equal(t, s.Meta["lettuce"], "7")
}

