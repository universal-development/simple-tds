package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	agent := r.UserAgent()
	redirect_id := vars["redirect_id"]
	fmt.Fprint(w, fmt.Sprintf("Redirect : %s for user agent: %s \n", redirect_id, agent))
}
