package grpc

import (
	"fmt"
	"net"

	"github.com/friendsofgo/wishlist/internal/api/grpc"

	"github.com/friendsofgo/wishlist/internal/listing"

	"github.com/friendsofgo/wishlist/internal/adding"
	"github.com/friendsofgo/wishlist/internal/creating"

	googlegrpc "google.golang.org/grpc"

	"github.com/friendsofgo/wishlist/internal/server"
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

	serviceServer := NewWishListServer(
		s.creatingService,
		s.addingService,
		s.listingService,
	)
	grpc.RegisterWishListServiceServer(googlegrpc.NewServer(), serviceServer)

	if err := googlegrpc.NewServer().Serve(listener); err != nil {
		return err
	}

	return nil
}
