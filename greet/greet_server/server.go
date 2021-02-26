package main

import (
	"com.grpc.tleu/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}

//LongGreet is an example of stream from client side
func (s *Server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")

	var result float32
	var count =float32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var res float32
			res = result / count
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		var x = float32(req.Greeting.GetNumber())
		result += x
		count++
	}
}

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
