package main

import (
	"fmt"
	"github.com/aaasen/gotwilio"
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

	twilio, err := gotwilio.NewTwilioClientFromFile("twilio-credentials.json")

	if err != nil {
		fmt.Println(err)
	}

	from := r.FormValue("To")
	to := r.FormValue("From")
	message := string(ddgResponse.GetAnswerText())
	twilio.SendSMS(from, to, message, "", "")

	fmt.Fprintf(w, string(ddgResponse.GetAnswerText()))
}
