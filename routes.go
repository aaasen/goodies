package main

import (
	"github.com/aaasen/dingo/dingo"
)

var routes = []*dingo.AHandler{
	dingo.NewHandler("GET", "/", QueryController{}),
}
