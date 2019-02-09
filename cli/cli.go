package main

import (
	"flag"

	"github.com/kristofferostlund/pfc-iban-validator/server"
)

var address = flag.String(
	"address",
	"localhost",
	"The address to run the server on",
)

var port = flag.Int(
	"port",
	5000,
	"The port to run the server on",
)

func main() {
	flag.Parse()

	srv := server.New(*address, *port)
	srv.Serve()
}
