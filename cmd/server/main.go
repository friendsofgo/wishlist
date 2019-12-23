package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/friendsofgo/wishlist/internal/adding"
	"github.com/friendsofgo/wishlist/internal/creating"
	"github.com/friendsofgo/wishlist/internal/listing"
	"github.com/friendsofgo/wishlist/internal/server"
	"github.com/friendsofgo/wishlist/internal/server/grpc"
	"github.com/friendsofgo/wishlist/internal/server/http"
	"github.com/friendsofgo/wishlist/internal/storage/inmemory"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/sync/errgroup"
)

const (
	WishListServerProtocolDefault = "tcp"
	WishListServerHostDefault     = "localhost"
	WishListServerPortDefault     = "3333"
)

func main() {
	var (
		protocol = getEnv("WISHLIST_SERVER_PROTOCOL", WishListServerProtocolDefault)
		host     = getEnv("WISHLIST_SERVER_HOST", WishListServerHostDefault)
		port     = getEnv("WISHLIST_SERVER_PORT", WishListServerPortDefault)

		repo            = inmemory.NewInMemoryWishListRepository()
		creatingService = creating.NewService(repo)
		addingService   = adding.NewService(repo)
		listingService  = listing.NewService(repo)
	)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		srvCfg := server.Config{Protocol: protocol, Host: host, Port: port}
		srv := grpc.NewServer(srvCfg, creatingService, addingService, listingService)

		log.Printf("gRPC server running at %s://%s:%s ...\n", protocol, host, port)
		return srv.Serve()
	})
	g.Go(func() error {
		httpAddr := fmt.Sprintf(":%s", port)
		httpSrv := http.NewServer(httpAddr)

		log.Printf("HTTP server running at %s ...\n", httpAddr)
		return httpSrv.Serve(ctx)
	})

	log.Fatal(g.Wait())
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
