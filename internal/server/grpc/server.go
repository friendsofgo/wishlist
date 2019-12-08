package grpc

import (
	"fmt"
	"net"
	"os"

	wishgrpc "github.com/friendsofgo/wishlist/genproto/go"
	"github.com/friendsofgo/wishlist/internal/adding"
	"github.com/friendsofgo/wishlist/internal/creating"
	"github.com/friendsofgo/wishlist/internal/listing"
	"github.com/friendsofgo/wishlist/internal/server"
	"github.com/friendsofgo/wishlist/internal/server/grpc/interceptor"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	srv := grpc.NewServer(withUnaryInterceptor())
	serviceServer := NewWishListServer(
		s.creatingService,
		s.addingService,
		s.listingService,
	)
	wishgrpc.RegisterWishListServiceServer(srv, serviceServer)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}

func withUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.LoggingServerInterceptor,
		interceptor.AuthorizationServerInterceptor,
	))
}
