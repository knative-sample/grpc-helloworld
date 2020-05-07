package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/knative-sample/grpc-helloworld/proto"
	"google.golang.org/grpc"
)

type HelloService struct{}

const (
	Port = "8080"
)

var waitTime = 10

func main() {
	maxConcurrentHellos, _ := strconv.Atoi(os.Getenv("GRPC_MAX_CONCURRENT_STREAMS"))
	waitTime, _ = strconv.Atoi(os.Getenv("WAIT_TIME"))
	server := grpc.NewServer(grpc.MaxConcurrentStreams(uint32(maxConcurrentHellos)))
	pb.RegisterHelloServiceServer(server, &HelloService{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("net.Listen error: %s", err)
	}

	server.Serve(lis)
}

func (s *HelloService) Hello(r *pb.HelloRequest, stream pb.HelloService_HelloServer) error {
	return fmt.Errorf("error hello")
}
