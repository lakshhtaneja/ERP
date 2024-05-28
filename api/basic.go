package api

import (
	"encoding/json"
	"net/http"

	"github.com/lakshhtaneja/ERP/data"
)

func Encode(d data.Student, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(d)
}
