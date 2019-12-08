package grpc

import (
	"fmt"
	"net"

	grpc "github.com/friendsofgo/wishlist/genproto/go"
	"github.com/friendsofgo/wishlist/internal/adding"
	"github.com/friendsofgo/wishlist/internal/creating"
	"github.com/friendsofgo/wishlist/internal/listing"
	"github.com/friendsofgo/wishlist/internal/server"
	googlegrpc "google.golang.org/grpc"
)

type grpcServer struct {
	config          server.Config
	creatingService creating.Service
	addingService   adding.Service
	listingService  listing.Service
}

func NewServer(
	config server.Config,
	cS creating.Service,
	aS adding.Service,
	lS listing.Service,
) server.Server {
	return &grpcServer{config: config, creatingService: cS, addingService: aS, listingService: lS}
}

func (s *grpcServer) Serve() error {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}

	srv := googlegrpc.NewServer()
	serviceServer := NewWishListServer(
		s.creatingService,
		s.addingService,
		s.listingService,
	)
	grpc.RegisterWishListServiceServer(srv, serviceServer)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}
