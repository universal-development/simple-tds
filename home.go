package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	agent := r.UserAgent()
	fmt.Fprint(w, fmt.Sprintf("Home handler, user agent: %s", agent))
}
