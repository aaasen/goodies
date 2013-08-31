package main

import (
	"github.com/aaasen/dingo"
)

var routes = []*dingo.AHandler{
	dingo.NewHandler("POST", "/", QueryController{}),
}
