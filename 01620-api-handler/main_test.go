package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFrameworkRouteHandler(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/frameworks", nil)
	FrameworkRouteHandler(recorder, req)
	if(recorder.Code != http.StatusOK) {
		t.Errorf("Wrong status, got %v expected %v", recorder.Code, http.StatusOK)
	}
}