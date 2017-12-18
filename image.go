package zvmsdk


import (
	"fmt"
	"bytes"
)


func ImageCreate() {
	var s []byte = make([]byte, 1)

	res := post("http://localhost:8080", s)
	fmt.Println("output is ", string(res))
}

func ImageDelete() {
	res := delete("http://localhost:8080", nil)
	fmt.Println("output is ", string(res))
}

func ImageGet(name string) {
	var buffer bytes.Buffer

	buffer.WriteString("http://localhost:8080/images")

	if name != "" {
		buffer.WriteString("?imagename=")
		buffer.WriteString(name)
	}
	res := get(buffer.String())

	fmt.Println("output is ", string(res))
}

