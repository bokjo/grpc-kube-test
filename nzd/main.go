package main

import (
	"log"
	"net"

	"github.com/bokjo/grpc-kube-test/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	nzd.RegisterNZDServiceServer(srv, &server{})

	reflection.Register(srv)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) Compute(ctx context.Context, r *nzd.NZDRequest) (*nzd.NZDResponse, error) {
	a, b := r.A, r.B

	for b != 0 {
		a, b = b, a%b
	}

	return &nzd.NZDResponse{Result: a}, nil
}
