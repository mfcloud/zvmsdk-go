package zvmsdk


import (
	"fmt"
	"bytes"
)


func ImageCreate() {

	res, result := post("http://localhost:8080", nil)
	fmt.Println("output is ", res, string(result))
}

func ImageDelete() {
	res, result := delete("http://localhost:8080", nil)
	fmt.Println("output is ", res, string(result))
}

func ImageGet(name string) {
	var buffer bytes.Buffer

	buffer.WriteString("http://localhost:8080/images")

	if name != "" {
		buffer.WriteString("?imagename=")
		buffer.WriteString(name)
	}
	res, result := get(buffer.String())

	fmt.Println("output is ", res, string(result))
}


