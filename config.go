package main

import (
	"github.com/aaasen/dingo"
)

var config = dingo.Config{
	Port:        "20480",
	TemplateDir: "/",
	StaticDir:   "/",
	Routes:      routes,
}
