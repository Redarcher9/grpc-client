package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Redarcher9/grpc-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8081"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	callSayHello(client)

}

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, _ := client.SayHello(ctx, &pb.NoParam{})
	res1, err := client.GetMobilesByOs(ctx, &pb.MobilesByOsRequest{
		Os: "IOS",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("%s", res.Message)
	log.Printf("%s", res1.Mobiles)
}
