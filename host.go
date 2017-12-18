package zvmsdk


import (
	"fmt"
)


func HostInfo() {
	resp, result := get("http://localhost:8080/abc")
	fmt.Println("output is ", resp, string(result))
}

func HostDiskpoolInfo() {
	resp, result := get("http://localhost:8080/abc1")
	fmt.Println("output is ", resp, string(result))
}
