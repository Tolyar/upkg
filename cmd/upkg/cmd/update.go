/*
Copyright © 2022 Tolyar
*/
package cmd

import (
	"github.com/Tolyar/upkg/pkg/packages"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command.
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update indexes of packages",
	Args:    cobra.MaximumNArgs(0),
	Aliases: []string{"makecache"},
	Run: func(cmd *cobra.Command, args []string) {
		p, err := packages.GetProvider()
		cobra.CheckErr(err)
		p.UpdateIndex()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
