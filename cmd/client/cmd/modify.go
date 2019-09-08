package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify allow change the name of your wish list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modify called")
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}
