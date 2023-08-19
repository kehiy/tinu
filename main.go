package main

import (
	model "github.com/kehiy/tinu/models"
	server "github.com/kehiy/tinu/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
