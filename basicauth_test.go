package basicauth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoAuthGetsDenied(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call handler")
	})
	handler := New("testRealm", map[string]string{"admin": "adminpass"}, []string{"GET"}, true)(nextHandler)

	w := &httptest.ResponseRecorder{}
	r, _ := http.NewRequest("GET", "/", nil)
	handler.ServeHTTP(w, r)

	assertDenied(t, w)
}

func TestCorrectCredentialsGetsAllowed(t *testing.T) {
	called := false
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	})

	handler := New("testRealm", map[string]string{"admin": "adminpass"}, []string{"GET"}, true)(nextHandler)

	w := &httptest.ResponseRecorder{}
	r, _ := http.NewRequest("GET", "/", nil)
	r.SetBasicAuth("admin", "adminpass")
	handler.ServeHTTP(w, r)

	assert.Equal(t, true, called)
	assertNotDenied(t, w)
}

func TestInvalidPasswordIsDeined(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call handler")
	})

	handler := New("testRealm", map[string]string{"admin": "adminpass"}, []string{"GET"}, true)(nextHandler)

	w := &httptest.ResponseRecorder{}
	r, _ := http.NewRequest("GET", "/", nil)
	r.SetBasicAuth("admin", "notadminspassword")
	handler.ServeHTTP(w, r)

	assertDenied(t, w)
}

func TestInvalidUserIsDenied(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should not call handler")
	})

	handler := New("testRealm", map[string]string{"admin": "adminpass"}, []string{"GET"}, true)(nextHandler)

	w := &httptest.ResponseRecorder{}
	r, _ := http.NewRequest("GET", "/", nil)
	r.SetBasicAuth("notadmin", "adminspassword")
	handler.ServeHTTP(w, r)

	assertDenied(t, w)
}

func assertNotDenied(t *testing.T, w *httptest.ResponseRecorder) {
	assert.NotEqual(t, http.StatusUnauthorized, w.Code)
}

func assertDenied(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, `Basic realm="testRealm"`, w.HeaderMap.Get("WWW-Authenticate"))
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
