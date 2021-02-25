package main

import (
	//"context"
	"fmt"
	"github.com/Arailym2002talgatkyzu/end-term-practice/greet/greetpb"
	"io"
	"log"
	"net"
	//"strconv"
	//"time"
	"google.golang.org/grpc"
)

//Server with embedded UnimplementedGreetServiceServer
type Server struct {
	greetpb.UnimplementedGreetServiceServer
}

//Greet is an example of unary rpc call

//LongGreet is an example of stream from client side
func (s *Server) AverageGreet(stream greetpb.GreetService_AverageGreetServer) error {
	fmt.Printf("AverageGreet function was invoked with a streaming request\n")
	var result int
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var res float32
			// we have finished reading the client stream
			res = float32(result / count)
			return stream.SendAndClose(&greetpb.AverageGreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		number := int(req.Greeting.GetNumber())
		result += number
		count++

	}
}

//GreetEveryone is an example of bidirectional stream

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
