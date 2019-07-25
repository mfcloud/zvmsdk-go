package zvmsdk

import (
        "net/http"
	"strings"
	"github.com/stretchr/testify/mock"
)


type HttpRequestMock struct {
        mock.Mock
}

func (h *HttpRequestMock) Get(url string) (int, []byte) {
        return 200, ""
}

func (h *HttpRequestMock) Post(url string, body []byte) (int, []byte) {
	return 200, ""
}

func (h *HttpRequestMock) Put(url string, body []byte, context RequestContext) (int, []byte) {
	return 200, ""
}

func (h *HttpRequestMock) Delete(url string, body []byte) (int, []byte){
	return 200, ""
}
