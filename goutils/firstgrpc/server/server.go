package main

import (
	"context"
	"log"
	"net"

	"github.com/diwangtseb/allutils/goutils/firstgrpc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAdminRpcServiceServer
}

func (s *server) GetAdminInfo(ctx context.Context, queryId *pb.QueryId) (*pb.Admin, error) {
	return &pb.Admin{
		Id:   1,
		Name: "gaoshou",
	}, nil
}

func main() {
	service, _ := net.Listen("tcp", ":5000")
	s := grpc.NewServer()
	pb.RegisterAdminRpcServiceServer(s, &server{})
	log.Printf("server listening at %v", service.Addr())
	s.Serve(service)
}
