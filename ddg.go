package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const DDG_URL = "http://api.duckduckgo.com/"

type DDGResponse struct {
	response string

	answer string
}

func NewDDGResponse(response string) *DDGResponse {
	return &DDGResponse{response, response}
}

func (response *DDGResponse) GetAnswer() string {
	return response.answer
}

type DDG struct {
	baseURL string
}

func QueryDDG(query string) *DDGResponse {
	queryURL := addArgToURL(DDG_URL, "q", query)
	queryURL = addArgToURL(queryURL, "format", "json")
	resp, err := http.Get(queryURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return NewDDGResponse(string(body))
}

func addArgToURL(baseURL string, key string, value string) string {
	u, err := url.Parse(baseURL)

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set(url.QueryEscape(key), url.QueryEscape(value))
	u.RawQuery = q.Encode()

	return u.String()
}
