package cmd

import (
	"fmt"

	wishgrpc "github.com/friendsofgo/wishlist/genproto/go"
	"github.com/spf13/cobra"
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

		w := &wishgrpc.WishList{
			Name:   name,
			Status: wishgrpc.WishList_ACTIVE,
		}

		res, err := cli.Create(ctx, &wishgrpc.CreateWishListReq{WishList: w})
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
