/*
Copyright © 2022 Tolyar
*/
package cmd

import (
	"github.com/Tolyar/upkg/pkg/packages"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command.
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show info for package",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := packages.GetProvider()
		cobra.CheckErr(err)
		p.Info(args...)
	},
	Aliases: []string{"show"},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
