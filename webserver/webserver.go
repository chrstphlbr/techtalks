package webserver

import (
	"net/http"

	"github.com/chrstphlbr/techtalks/controller"
	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/videos", controller.DummyFunc)
	http.Handle("/", r)
}
