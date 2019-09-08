package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	googlegrpc "google.golang.org/grpc"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/friendsofgo/wishlist/internal/net/grpc"
)

var cfgFile string
var cli grpc.WishListServiceClient
var ctx context.Context

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wishlist",
	Short: "A simple gRPC wish list, to add items to your wish lists",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// initConfig reads in config file and ENV variables
	cobra.OnInitialize(initConfig)

	var (
		host    = os.Getenv("WISHLIST_HOST_SERVER")
		port, _ = strconv.Atoi(os.Getenv("WISHLIST_PORT_SERVER"))
	)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())

	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}

	cli = grpc.NewWishListServiceClient(conn)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
