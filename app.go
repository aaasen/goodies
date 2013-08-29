package main

import (
	"github.com/aaasen/dingo/dingo"
)

func main() {
	server := dingo.New(config)
	server.Run()
}
