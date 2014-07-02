package controller

import (
	"net/http"

	"github.com/chrstphlbr/techtalks/data"
)

func init() {

}

func DummyFunc(res http.ResponseWriter, req *http.Request) {
	fi := data.NewFileImporter("/Users/Christoph/tmp/techtalks_data/")
	_ = fi
}
