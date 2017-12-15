package zvmsdk


import (
	"fmt"
)


func HostInfo() {
	res := get("http://localhost:8080/abc")
	fmt.Println("output is ", string(res))
}

func HostDiskpoolInfo() {
	res := get("http://localhost:8080/abc1")
	fmt.Println("output is ", string(res))
}


