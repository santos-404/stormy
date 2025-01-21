/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stormy",
	Short: "A simple CLI tool to manage all your passwords. Built with Go and Cobra.",
	Long: `Stormy is a command-line tool built with Go and Cobra to help you manage 
and organize your passwords securely and efficiently. 

Before saving your passwords, you must start setting a master password.
You can do this with:
	stormy set-master-password <your-password>

Features include:
- Adding, retrieving, and deleting passwords.
- Secure storage and encryption.
- Easy-to-use commands for quick access.

Get started by running 'stormy help' for a list of available commands.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passwordmanager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
