package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	var (
		host    = os.Getenv("WISHLIST_HOST_SERVER")
		port, _ = strconv.Atoi(os.Getenv("WISHLIST_PORT_SERVER"))
	)

	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}

}
