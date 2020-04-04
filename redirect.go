package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	agent := r.UserAgent()
	redirectId := vars["redirectId"]
	url := resolve(redirectId, agent)

	fmt.Fprint(w, fmt.Sprintf("Redirect id : %s agent : %s redirect url: %s", redirectId, agent, url))
}
