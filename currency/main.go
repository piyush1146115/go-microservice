package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/piyush1146115/go-microservice/currency/data"
	protos "github.com/piyush1146115/go-microservice/currency/protos/currency"
	"github.com/piyush1146115/go-microservice/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()

	log.Info("Starting currency service", "port", 9092)

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	c := server.NewCurrency(rates, log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, c)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	gs.Serve(l)
}
