package cmd

import (
	"fmt"

	wishgrpc "github.com/friendsofgo/wishlist/genproto/go"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the items of the given wish list",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		res, err := cli.List(ctx, &wishgrpc.ListWishListReq{WishListId: id})
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	listCmd.Flags().StringP("id", "i", "", "Id of wish list")
	_ = listCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(listCmd)
}
