/*
Copyright © 2022 Gabriel de Lisle
*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)

var datafile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo list application",
	Long:  "Tri will help you manage your tasks more efficiently",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gabriel-de-lisle.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Unable to find root directory")
	}
	rootCmd.Flags().StringVar(&datafile, "datafile", home+string(os.PathSeparator)+".tri.json", "File to save tasks")
}
