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
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(ddgResponse.GetAnswer()))
}
