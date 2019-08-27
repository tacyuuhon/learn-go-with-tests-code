package server

import (
	"fmt"
	"net/http"
)

// PlayerServer func
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
