package main

import (
	"encoding/json"
	"fmt"
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
	Height interface{}
	Width  interface{}
}

func NewDDGResponse(response []byte) (*DDGResponse, error) {
	var m DDGResponse
	err := json.Unmarshal(response, &m)

	fmt.Println(string(response))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	m.Answer = m.getAnswer()

	return &m, nil
}

// Generates a little summary of the DDG response.
// For example, the summary of a definition would include only the definition,
// and the summary of a calculation would include only the answer
func (response *DDGResponse) Summarize() string {
	if len(response.Definition) > 0 {
		return response.Definition
	} else if len(response.Answer) > 0 {
		return response.Answer
	} else if len(response.Abstract) > 0 {
		return response.Abstract
	} else {
		return "=("
	}
}

func (response *DDGResponse) getAnswer() string {
	xmlTagRegex := regexp.MustCompile("<.*?>")
	answerText := xmlTagRegex.ReplaceAllString(response.Answer, "")

	return answerText
}
