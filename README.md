basicauth-go
=================
[![GoDoc](https://godoc.org/github.com/99designs/basicauth-go?status.svg)](https://godoc.org/github.com/99designs/basicauth-go)
[![Build Status](https://travis-ci.org/99designs/basicauth-go.svg)](https://travis-ci.org/99designs/basicauth-go)


golang middleware for HTTP basic auth.

```go
// Chi

router.Use(basicauth.New("testRealm", map[string]string{"admin": "adminpass"}, []string{"GET"}, true))


// Manual wrapping

middleware := basicauth.New("testRealm", map[string]string{"admin": "adminpass"}, []string{"GET"}, true)

h := middlware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {
    /// do stuff
})

log.Fatal(http.ListenAndServe(":8080", h))
```

### env loading
If your environment looks like this:
```bash
SOME_PREFIX_BOB=password
SOME_PREFIX_JANE=password
```

you can load it like this:
```go
middleware := basicauth.NewFromEnv("MyRealm", "SOME_PREFIX", []string{"GET"}, true)
```

