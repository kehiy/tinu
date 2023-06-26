package main

import (
	"github.com/kehiy/tinu/models"
	"github.com/kehiy/tinu/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
