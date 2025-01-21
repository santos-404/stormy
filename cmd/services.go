/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/stormy/utils"
	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "A list of every service stored",
	Long: `This command provides a list of every service stored in the password manager.

Examples:
	# List all the services stored
	stormy services
	`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ListAllServices()
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}
