package util

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseWithJson(t *testing.T) {
	//r := httptest.NewRequest("GET", "/movie", nil)
	w := httptest.NewRecorder()

	ResponseWithJson(w, http.StatusOK, "")
}
