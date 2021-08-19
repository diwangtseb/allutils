package main

import (
	"context"
	"log"
	"time"

	"github.com/diwangtseb/allutils/goutils/firstgrpc/pb"

	"google.golang.org/grpc"
)

const (
	addr = "localhost:5000"
)

func main() {
	conn, _ := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	c := pb.NewAdminRpcServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.GetAdminInfo(ctx, &pb.QueryId{Id: 1})
	log.Println(r.GetName())
}
