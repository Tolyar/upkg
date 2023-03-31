/*
Copyright © 2022 none

*/
package cmd

import (
	"github.com/Tolyar/upkg/pkg/packages"
	"github.com/spf13/cobra"
)

// providesCmd represents the provides command
var providesCmd = &cobra.Command{
	Use:   "provides",
	Short: "Show package wich provides resource.",
	Long: `Show package wich provides resource.
	Not all providers supports this option.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := packages.GetProvider()
		cobra.CheckErr(err)
		p.Search(args...)
	},
	Aliases: []string{"which"},
}

func init() {
	rootCmd.AddCommand(providesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// providesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// providesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
