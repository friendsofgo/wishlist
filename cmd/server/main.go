package main

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"

	"github.com/friendsofgo/wishlist/internal/server/grpc"
	"github.com/friendsofgo/wishlist/internal/server/grpc/handler"
)

func main() {
	var (
		protocol = os.Getenv("WISHLIST_PROTOCOL_SERVER")
		host     = os.Getenv("WISHLIST_HOST_SERVER")
		port, _  = strconv.Atoi(os.Getenv("WISHLIST_PORT_SERVER"))
	)

	// gRPC services
	var (
		wishListService = handler.NewWishList()
	)

	srv := grpc.NewServer(protocol, host, port, wishListService)

	log.Printf("gRPC server running at %s://%s:%d ...\n", protocol, host, port)
	log.Fatal(srv.Run())
}
