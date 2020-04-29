package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	pb "github.com/knative-sample/grpc-helloworld/proto"
	"google.golang.org/grpc"
)

func main() {

	countStr := os.Getenv("GRPC_CONCURRENT")
	serverAddr := os.Getenv("GRPC_SERVER_ADDR")
	conn, err := grpc.Dial(
		serverAddr,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			//grpc.WithTimeout(time.Second * 300),
		}...,
	)
	if err != nil {
		log.Fatalf("grpc.Dial error: %s", err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	for {
		count, _ := strconv.Atoi(countStr)
		if count == 0 {
			count = 1000
		}
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func(index int) {
				err = printHelloLists(
					client,
					&pb.HelloRequest{
						Msg: &pb.HelloMessage{
							Key:   fmt.Sprintf("Client Hello for [%d]", index),
							Value: int32(index),
						},
					},
				)
				if err != nil {
					log.Printf("printHello error: %s", err)
				}
				wg.Done()
			}(i)
		}

		wg.Wait()
	}

}

func printHelloLists(client pb.HelloServiceClient, r *pb.HelloRequest) error {
	stream, err := client.Hello(context.Background(), r)
	if err != nil {
		return err
	}

	stime := time.Now()
	var resp *pb.HelloResponse
	for {
		_resp, err := stream.Recv()
		if err != nil {
			return err
		}
		resp = _resp
		log.Printf("resp: key: %s, latency:%d", resp.Msg.Key, time.Now().Unix()-stime.Unix())
		if resp.Msg.Key == "last" {
			break
		}
	}

	stream.CloseSend()
	return nil
}
