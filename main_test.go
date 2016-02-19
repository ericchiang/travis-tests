package main

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHostname(t *testing.T) {

	host := "travis.dev"

	gotRequest := false

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if reqHost, _, err := net.SplitHostPort(r.Host); err == nil {
			gotRequest = reqHost == host
		}

		t.Log("Received request for host", r.Host)
		w.WriteHeader(http.StatusOK)
	}))
	defer s.Close()

	defer func() {
		if !gotRequest {
			t.Error("Did not get request for host", host)
		}
	}()

	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	_, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get("http://" + host + ":" + port + "/")
	if err != nil {
		t.Error("GET failed", err)
		return
	}
	resp.Body.Close()
}
