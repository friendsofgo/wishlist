package http

import (
	"context"
	"net/http"

	wishgrpc "github.com/friendsofgo/wishlist/genproto/go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	httpAddr string
}

func NewServer(httpAddr string) *Server {
	return &Server{httpAddr: httpAddr}
}

func (s *Server) Serve(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := wishgrpc.RegisterWishListServiceHandlerFromEndpoint(ctx, mux, s.httpAddr, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(s.httpAddr, mux)
}
