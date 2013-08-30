package main

import (
	"fmt"
	"net/http"
)

type QueryController struct{}

func (c QueryController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	query := r.FormValue("q")

	ddgResponse, err := QueryDDG(query)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprintf(w, string(ddgResponse.Answer))
}
