package api

import (
	"net/http"

	"github.com/lakshhtaneja/ERP/internal"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	Encode(internal.Peter, w, r)
}
