package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPHandler_404(t *testing.T) {
	handler := NewHTTPRouter()
	request := httptest.NewRequest(http.MethodGet, "/nope?a=1&b=2", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusNotFound {
		t.Error("want 404, got:", response.Code)
	}
}
func TestHTTPHandler_200(t *testing.T) {
	testHTTP200(t, "/add?a=1&b=2", "3")
	testHTTP200(t, "/sub?a=5&b=3", "2")
	testHTTP200(t, "/mul?a=3&b=4", "12")
	testHTTP200(t, "/div?a=100&b=50", "2")
	testHTTP200(t, "/bog?a=1&b=2", "45")
}
func testHTTP200(t *testing.T, path, expectedResponseBody string) {
	handler := NewHTTPRouter()
	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Error("non-200 status:", response.Code)
	}
	actualResponseBody := strings.TrimSpace(response.Body.String())
	if actualResponseBody != expectedResponseBody {
		t.Errorf("want '%s', got: '%s'", expectedResponseBody, actualResponseBody)
	}
}
func TestHTTPHandler_422_InvalidArgA(t *testing.T) {
	handler := NewHTTPRouter()
	request := httptest.NewRequest(http.MethodGet, "/add?a=NaN&b=2", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusUnprocessableEntity {
		t.Error("want 422, got:", response.Code)
	}
	responseBody := strings.TrimSpace(response.Body.String())
	if responseBody != "invalid 'a' parameter: [NaN]" {
		t.Error("bad response body:", responseBody)
	}
}
func TestHTTPHandler_422_InvalidArgB(t *testing.T) {
	handler := NewHTTPRouter()
	request := httptest.NewRequest(http.MethodGet, "/add?a=1&b=NaN", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusUnprocessableEntity {
		t.Error("want 422, got:", response.Code)
	}
	responseBody := strings.TrimSpace(response.Body.String())
	if responseBody != "invalid 'b' parameter: [NaN]" {
		t.Error("bad response body:", responseBody)
	}
}
