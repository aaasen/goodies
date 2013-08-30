package main

import (
	"fmt"
	"net/http"
)

type QueryController struct{}

func (c QueryController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	query := r.FormValue("q")

	fmt.Fprintf(w, QueryDDG(query).GetAnswer())
}
