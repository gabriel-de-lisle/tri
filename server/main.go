package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/gabriel-de-lisle/tri/protocol"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 5001, "The server port")
	datafile = flag.String("datafile", ".tri.json", "The file to save tasks")
)

type server struct {
	pb.UnimplementedTaskHandlerServer
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskHandlerServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
