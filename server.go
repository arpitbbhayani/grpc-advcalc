package main

import (
	"log"
	"net"

	"github.com/arpitbbhayani/grpc-calc/advcalc"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("unable to listen on 8081. Port already in use")
	}

	grpcServer := grpc.NewServer()

	aServer := advcalc.AdvancedCalcServerLive{}
	advcalc.RegisterAdvancedCalcServer(grpcServer, &aServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal((err))
	}
}
