package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/Arailym2002talgatkyzu/end-term-practice/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doAverageGreet(c)
}


func doAverageGreet(c greetpb.GreetServiceClient) {

	requests := []*greetpb.AverageGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				Number: 1,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: 2,
			},
		},
		{
			Greeting: &greetpb.Greeting{
				Number: 3,
			},
		},
	}

	ctx := context.Background()
	stream, err := c.AverageGreet(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}




