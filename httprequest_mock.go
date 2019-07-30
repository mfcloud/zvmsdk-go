package zvmsdk

import (
	"github.com/stretchr/testify/mock"
)

type HttpRequestMock struct {
	mock.Mock
}

func (h *HttpRequestMock) Get(url string) (int, []byte) {
	return 200, []byte("")
}

func (h *HttpRequestMock) Post(url string, body []byte) (int, []byte) {
	ret := h.Called(url, body)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, []byte) int); ok {
		r0 = rf(url, body)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string, []byte) []byte); ok {
		r1 = rf(url, body)
	} else {
		r1 = ret.Get(1).([]byte)
	}

	return r0, r1
}

func (h *HttpRequestMock) Put(url string, body []byte, context RequestContext) (int, []byte) {
	return 200, []byte("")
}

func (h *HttpRequestMock) Delete(url string, body []byte) (int, []byte) {
	return 200, []byte("")
}
