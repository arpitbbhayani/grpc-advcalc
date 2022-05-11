package main

import (
	"context"
	"log"

	"github.com/arpitbbhayani/grpc-calc/advcalc"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	a := advcalc.NewAdvancedCalcClient(conn)
	res, err := a.Power(context.Background(), &advcalc.PowerRequest{
		A: 10,
		B: 2,
	})
	if err != nil {
		log.Print(err)
	} else {
		log.Print(res.Pow)
	}
}
