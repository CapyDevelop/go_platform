package hello

import (
	"context"
	pb "github.com/CapyDevelop/go_platform/skeleton/api/gen/go/hello"
	"google.golang.org/grpc"
)

type HelloWorld interface {
	SayHello(ctx context.Context)
}

type serverAPI struct {
	pb.UnimplementedHelloWorldServer
}

func Register(gRPC *grpc.Server) {
	pb.RegisterHelloWorldServer(gRPC, &serverAPI{})
}

func (s *serverAPI) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Answer: "hello" + request.Name}, nil
}

func (s *serverAPI) SaveHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Answer: "hello" + request.Name}, nil
}
