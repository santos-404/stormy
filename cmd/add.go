/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/stormy/utils"
	"github.com/spf13/cobra"
)

var (
	service  string
	username string
	password string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password entry",
	Long: `Add a new password entry for a specific 
service with a username and password.

Examples:
	# Add a new password entry for the service 'github' with username 'johndoe' and password 'password123'
	stormy add -s github -u johndoe -p password123`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.AddPassword(service, username, password)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&service, "service", "s", "", "Name of the service (required)")
	addCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the service (required)")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "Password for the service (required)")

	addCmd.MarkFlagRequired("service")
	addCmd.MarkFlagRequired("username")
	addCmd.MarkFlagRequired("password")

}
