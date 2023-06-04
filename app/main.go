package main

import (
	"fmt"

	"github.com/kehiy/tinu/model"
	"github.com/kehiy/tinu/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}