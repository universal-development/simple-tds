package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	agent := r.UserAgent()
	redirectID := vars["redirectId"]
	url := resolve(redirectID, agent)

	fmt.Fprint(w, fmt.Sprintf("Redirect id : %s agent : %s redirect url: %s", redirectID, agent, url))
}
