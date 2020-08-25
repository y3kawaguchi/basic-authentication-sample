package main

import (
	"fmt"
	"net/http"
)

const (
	basicAuthUser     = "user"
	basicAuthPassword = "pass"
)

func main() {
	http.HandleFunc("/basicAuth",
		func(w http.ResponseWriter, r *http.Request) {
			if user, pass, ok := r.BasicAuth(); !ok || user != basicAuthUser || pass != basicAuthPassword {
				w.Header().Add("WWW-Authenticate", `Basic realm="my private area"`)
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Unauthorized", 401)
				return
			}
			fmt.Fprintf(w, "Authorized")
		},
	)
	http.ListenAndServe(":18888", nil)
}
