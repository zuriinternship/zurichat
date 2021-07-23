package data

import (
	"fmt"
	"net/http"
)

func ReadData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is you reading data\n")
}
