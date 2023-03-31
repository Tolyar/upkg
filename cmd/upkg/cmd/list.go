/*
Copyright © 2022 Tolyar
*/
package cmd

import (
	"github.com/Tolyar/upkg/pkg/packages"
	"github.com/spf13/cobra"
)

var listUpgradable = false

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed packages",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := packages.GetProvider()
		cobra.CheckErr(err)
		if listUpgradable {
			p.ListUpdates()
		} else {
			p.ListInstalled()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolVarP(&listUpgradable, "updates", "u", false, "List upgradable packages only")
}
