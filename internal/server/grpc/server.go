package grpc

import (
	"fmt"
	"net"

	googlegrpc "google.golang.org/grpc"

	"github.com/friendsofgo/wishlist/internal/net/grpc"
	"github.com/friendsofgo/wishlist/internal/server"
)

type grpcServer struct {
	protocol        string
	host            string
	port            int
	wishListHandler grpc.WishListServiceServer
}

func NewServer(
	protocol string,
	host string,
	port int,
	wishListHandler grpc.WishListServiceServer,
) server.Server {
	return &grpcServer{
		protocol: protocol,
		host:     host,
		port:     port,

		wishListHandler: wishListHandler,
	}
}

func (s *grpcServer) Run() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	listener, err := net.Listen(s.protocol, addr)
	if err != nil {
		return err
	}

	srv := googlegrpc.NewServer()
	grpc.RegisterWishListServiceServer(srv, s.wishListHandler)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}
