/*
Copyright © 2022 Tolyar
*/
package cmd

import (
	"github.com/Tolyar/upkg/pkg/packages"
	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command.
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade packages",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := packages.GetProvider()
		cobra.CheckErr(err)
		if len(args) == 0 {
			p.UpgradeAll()
		} else {
			p.Upgrade(args...)
		}
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upgradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
