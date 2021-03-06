package main

import (
	"github.com/dejavuzhou/evefs/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":12345"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRpcServiceServer(s, &rpcServer{})
	s.Serve(lis)
}
