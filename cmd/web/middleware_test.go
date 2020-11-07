package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want %q; got %q", "deny", frameOptions)
	}

	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("want %q; got %q", "1; mode=block", xssProtection)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

// func newTestApplication(t *testing.T) *application {
// 	return &application{
// 		errorLog: log.New(ioutil.Discard, "", 0),
// 		infoLog:  log.New(ioutil.Discard, "", 0),
// 	}
// }

func TestRequiredAuthentication(t *testing.T) {
	user := "testuser2"
	email := "testuser2@gmail.com"
	password := "Test@123"

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	app.users.Insert(user, email, password)

	// code, _, body := ts.get(t, "/user/login")

	// if code != http.StatusSeeOther {
	// 	t.Errorf("want %d; got %d", http.StatusSeeOther, code)
	// }

	// if string(body) != "StatusSeeOther" {
	// 	t.Errorf("want body to equal %q", "StatusSeeOther")
	// }

	// rr := httptest.NewRecorder()
	// err := app.users.Insert(user, email, password)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("OK"))
	// })

	// ts := httptest.NewTLSServer(app.routes())
	// defer ts.Close()

	// r, err := http.NewRequest(http.MethodGet, "/snippet/create", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// id, err := app.users.Authenticate(email, password)
	// app.session.Put(r, "authenticatedUserID", id)
	// app.authenticate(next)
	// app.requireAuthentication(next).ServeHTTP(rr, r)

	// rs := rr.Result()

	// cacheControl := rs.Header.Get("Cache-Control")
	// if cacheControl != "no-store" {
	// 	t.Errorf("want %q; got %q", "no-store", cacheControl)
	// }

	// if rs.StatusCode != http.StatusOK {
	// 	t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	// }

	// defer rs.Body.Close()

	// body, err := ioutil.ReadAll(rs.Body)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if string(body) != "OK" {
	// 	t.Errorf("want body to equal %q", "OK")
	// }
}
