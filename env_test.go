package basicauth

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvLoading(t *testing.T) {
	os.Setenv("TESTAPI_BOB", "bobspassword")

	called := false
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	handler := NewFromEnv("testrealm", "TESTAPI_", []string{"GET"}, true)(nextHandler)

	w := &httptest.ResponseRecorder{}
	r, _ := http.NewRequest("GET", "/", nil)
	r.SetBasicAuth("bob", "bobspassword")
	handler.ServeHTTP(w, r)

	assert.Equal(t, true, called)
	assertNotDenied(t, w)
}
