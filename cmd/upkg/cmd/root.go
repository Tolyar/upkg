package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "upkg",
	Short: "Universal package manager",
	Long:  `upkg can be used instead of yum, dnf, apt, ...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			//nolint
			cmd.Help()
			os.Exit(0)
		}
		os.Exit(0)
		// pp.Println("OS:", sysinfo.OS())
		// pp.Println("Arch:", sysinfo.Arch())
		// if pm, err := packages.PMName(); err == nil {
		// 	pp.Println("package manager:", pm)
		// } else {
		// 	pp.Println("package manager is unknown.", err)
		// }
		// if sysinfo.OS() == "linux" {
		// 	lr, err := sysinfo.LinuxRelease()
		// 	if err != nil {
		// 		log.Fatalf("Error %+v", err)
		// 	}
		// 	pp.Println("Linux info:", lr)
		// }
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.upkg.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".upkg" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".upkg")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
