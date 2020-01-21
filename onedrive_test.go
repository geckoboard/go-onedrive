package onedrive

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var (
	mux      *http.ServeMux
	server   *httptest.Server
	oneDrive *OneDrive
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	oneDrive = New(http.DefaultClient)
	oneDrive.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func fileWrapperHandler(file string, status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		b, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	}
}
