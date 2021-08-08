package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/building-microservices-youtube/currency/server"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	grpcServer := grpc.NewServer()
	currencyServer := server.NewCurrency(log)

	protos.RegisterCurrencyServer(grpcServer, currencyServer)

	netListener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(netListener)
}
