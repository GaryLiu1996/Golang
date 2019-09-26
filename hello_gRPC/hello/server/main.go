package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "pro/app_pitt_bff_20190918/proto/hello"
)

const (
	Address = "127.0.0.1:50052"
)

//定义结构体用作实现约定接口
type helloService struct {

}

var HelloService = helloService{}

//SayHello函数实现了helloService接口
func (h helloService)SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloResponse,error){
	resp := new(pb.HelloResponse)

	resp.Message = fmt.Sprintf("Hello %s",in.Name)

	return resp,nil
}

func main(){
	listen,err:=net.Listen("tcp",Address)

	if err!=nil{
		grpclog.Fatalf("Failed to listen:%v",err)
	}

	s:=grpc.NewServer()

	pb.RegisterHelloServer(s,HelloService)

	fmt.Printf("Listen on : "+Address)

	s.Serve(listen)
}