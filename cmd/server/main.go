package main

import (
	"log"
	"os"

	"github.com/friendsofgo/wishlist/internal/adding"
	"github.com/friendsofgo/wishlist/internal/listing"

	"github.com/friendsofgo/wishlist/internal/storage/inmemory"

	"github.com/friendsofgo/wishlist/internal/creating"

	"github.com/friendsofgo/wishlist/internal/server"

	_ "github.com/joho/godotenv/autoload"

	"github.com/friendsofgo/wishlist/internal/server/grpc"
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

	srvCfg := server.Config{Protocol: protocol, Host: host, Port: port}
	srv := grpc.NewServer(srvCfg, creatingService, addingService, listingService)

	log.Printf("gRPC server running at %s://%s:%s ...\n", protocol, host, port)
	log.Fatal(srv.Serve())
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
