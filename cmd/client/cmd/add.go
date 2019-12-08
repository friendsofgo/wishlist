package cmd

import (
	"fmt"

	wishgrpc "github.com/friendsofgo/wishlist/genproto/go"
	"github.com/spf13/cobra"
)

// addCmd represents the modify command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new item to the given wish list",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		link, err := cmd.Flags().GetString("link")
		if err != nil {
			return err
		}

		price, err := cmd.Flags().GetFloat64("price")
		if err != nil {
			return err
		}

		i := &wishgrpc.Item{
			Name:       name,
			WishListId: id,
			Link:       link,
			Price:      price,
			Priority:   wishgrpc.Item_LOW,
			Status:     wishgrpc.Item_ACTIVE,
		}

		res, err := cli.Add(ctx, &wishgrpc.AddItemReq{Item: i})
		if err != nil {
			return err
		}

		fmt.Println(res)
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("id", "i", "", "Id of the wish list")
	_ = addCmd.MarkFlagRequired("id")

	addCmd.Flags().StringP("name", "n", "", "Name of the item")
	_ = addCmd.MarkFlagRequired("name")

	addCmd.Flags().StringP("link", "l", "", "Link of the item")
	_ = addCmd.MarkFlagRequired("link")

	addCmd.Flags().Float64P("price", "p", 0.0, "Price of the item")
	_ = addCmd.MarkFlagRequired("price")

	rootCmd.AddCommand(addCmd)
}
