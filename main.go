package main

import (
	"api/src/initializers"
	"api/src/presentation/server"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	server.InitServer()
}
