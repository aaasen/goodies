package main

import (
	"github.com/aaasen/dingo/dingo"
)

var config = dingo.Config{
	Port:        "8080",
	TemplateDir: "/",
	StaticDir:   "/",
	Routes:      routes,
}
