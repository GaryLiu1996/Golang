package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "pro/app_pitt_bff_20190918/proto/hello"
)

const (
	Address = "127.0.0.1:50052" +
		"" +
		""
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{Name: "Grpc!"}

	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}
	println(res.Message)
	//grpclog.Infoln(res.Message)
}
