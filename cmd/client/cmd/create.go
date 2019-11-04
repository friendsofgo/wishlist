package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	wishlist "github.com/friendsofgo/wishlist/internal"
	"github.com/friendsofgo/wishlist/internal/net/grpc"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Store a new wish list",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		w := &grpc.WishList{
			Id:     wishlist.UlidGen(),
			Name:   name,
			Status: grpc.WishList_ACTIVE,
		}

		res, err := cli.Create(ctx, &grpc.CreateWishListReq{WishList: w})
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of wish list")
	_ = createCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createCmd)
}
