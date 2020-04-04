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

	fmt.Fprint(w, fmt.Sprintf(`<!DOCTYPE html>
		<html>
		<head>
		   <meta http-equiv="refresh" content="0; URL='%s'" />
		   <script type="text/javascript">
		            window.location.href = "%s"
		        </script>
		</head>
		<body>

		</body>
		</html>

		`, url, url))
}
