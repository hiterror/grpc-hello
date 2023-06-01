package main

import (
	"context"
	pb "demogrpc/proto"
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "jiujiayi",
		"appKey": "123456",
	}, nil
}
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return true
}

func main() {
	creds, _ := credentials.NewClientTLSFromFile(
		"./configs/test.pem",
		"*.jiujiayi.com",
	)
	_ = creds

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		glog.Error(err)
		return
	}
	defer conn.Close()

	client := pb.NewSayHelloClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		RequestName: "blake",
	})
	if err != nil {
		glog.Error(err)
		return
	}

	fmt.Println(resp.GetResponseMsg())
}
