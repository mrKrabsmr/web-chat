package main

import (
	"flag"
	server "mrKrabsmr/web-chat/internal"

	"github.com/joho/godotenv"
)

var (
	version = flag.Int("v", 1, "version")
)

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	apiServer := server.NewAPIServer()
	apiServer.ConfigureRoutes(*version)
	apiServer.Run()
}
