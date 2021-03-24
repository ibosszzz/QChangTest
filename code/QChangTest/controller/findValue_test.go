package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)

func TestFindValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/findValue", bytes.NewReader([]byte{}))
	res := httptest.NewRecorder()

	FindValue(res, req)
	
	if res.Code != http.StatusOK {
		t.Error(res.Body.String())
	} else {
		t.Log(res.Body.String())
	}
}