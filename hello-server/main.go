package main

import (
	"context"
	pb "demogrpc/proto"
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
	pb.UnimplementedSayHelloServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("md not ok")
	}
	fmt.Println(md)

	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	name := req.GetRequestName()
	if len(name) == 0 {
		return nil, errors.New("name empty")
	}
	return &pb.HelloResponse{
		ResponseMsg: fmt.Sprintf("hello %s, appId %s, appKey %s", name, appId, appKey),
	}, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile(
		"/Users/caiwei/tech/terminate/example/grpc/configs/test.pem",
		"/Users/caiwei/tech/terminate/example/grpc/configs/test.key",
	)
	if err != nil {
		glog.Error(err)
		return
	}
	_ = creds

	// 开启端口
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		glog.Error(err)
		return
	}

	// 创建grpc服务
	//grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// 在grpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, NewServer())

	// 启动服务
	err = grpcServer.Serve(lis)
	if err != nil {
		glog.Error("serve error", err)
	}
}
