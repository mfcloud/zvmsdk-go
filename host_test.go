package zvmsdk

import (
	"testing"
	"github.com/stretchr/testify/require"

)


func Test_HostInfo(t *testing.T) {
        //FIXME: mock this later
        status, _ := HostInfo()
        require.Equal(t, status, 200)
}

func Test_HostDiskPoolInfo(t *testing.T) {
        //FIXME: mock this later
        status, _ := HostDiskpoolInfo("disk1")
        require.Equal(t, status, 200)
}
