package main

import (
	"fmt"
	"net/http"
)

type QueryController struct{}

func (c QueryController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	query := r.FormValue("Body")

	ddgResponse, err := QueryDDG(query)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprintf(w, string(ddgResponse.GetAnswerText()))
}
