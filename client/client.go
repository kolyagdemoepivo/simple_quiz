/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"client/api"
	"client/cmd"
	"context"
	"flag"
	"log"
	pb "quiz_schema"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:6690", "the address to connect to")
)

func main() {
	ctx := context.Background()

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewQuizGrpcClient(conn)
	quizApi := api.NewQuizApi(c, ctx)
	cmd.Execute(quizApi)
}
