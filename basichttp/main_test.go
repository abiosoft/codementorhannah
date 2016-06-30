package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/hello", handleHello)
	s.HandleFunc("/error", handleError)

	// simulate a server using an handler function
	server := httptest.NewServer(s)

	tests := []struct {
		endpoint   string
		statusCode int
		status     string
	}{
		{"/hello", http.StatusOK, "success"},
		{"/error", http.StatusInternalServerError, "error"},
		{"/notfound", http.StatusNotFound, ""},
	}

	t.Logf("Listening on %s", server.URL)
	for i, test := range tests {
		// making a GET request to the simulate server
		resp, err := http.Get(server.URL + test.endpoint)
		if err != nil {
			t.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		t.Logf("%v --> %d with body %v", test.endpoint, resp.StatusCode,
			string(body))

		// status code
		if resp.StatusCode != test.statusCode {
			t.Errorf("Test %d: Expected %d, found %d", i,
				test.statusCode, resp.StatusCode)
		}
	}

}

//
func TestMain(t *testing.M) {
	os.Exit(t.Run())
}
