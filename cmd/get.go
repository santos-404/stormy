/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/passwordmanager/storage"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password given a service and username",
	Long: `The get command retrieves a stored password for a specified service and username.
This command allows you to quickly access your saved passwords without having to manually look them up.

Examples:
  # Get the password for the service 'github' and username 'johndoe'
  passwordmanager get -s github -u johndoe

  # Get the password for the service 'email' and username 'janedoe'
  passwordmanager get -s email -u janedoe`,
	Run: func(cmd *cobra.Command, args []string) {
		storage.GetPassword(service, username)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&service, "service", "s", "", "Name of the service (required)")
	getCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the service (required)")

	getCmd.MarkFlagRequired("service")
	getCmd.MarkFlagRequired("username")

}
