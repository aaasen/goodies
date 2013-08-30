package main

import (
	"encoding/json"
	"regexp"
)

type DDGResponse struct {
	response []byte

	Definition       string
	DefinitionSource string
	DefinitionURL    string

	Abstract       string
	AbstractText   string
	AbstractSource string
	AbstractURL    string

	Answer     string
	AnswerType string

	Heading       string
	Image         string
	RelatedTopics []RelatedTopic
	Redirect      string
	Type          string
	Results       []Result
}

type Result struct {
	Result   string
	FirstURL string
	Text     string
}

type RelatedTopic struct {
	Result   string
	Icon     Icon
	FirstURL string
	Text     string
}

type Icon struct {
	URL    string
	Height string
	Width  string
}

func NewDDGResponse(response []byte) (*DDGResponse, error) {
	var m DDGResponse
	err := json.Unmarshal(response, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (response *DDGResponse) GetAnswerText() string {
	xmlTagRegex := regexp.MustCompile("<.*?>")
	answerText := xmlTagRegex.ReplaceAllString(response.Answer, "")

	return answerText
}
