/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/passwordmanager/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list of every password stored (this does not list the passwords themselves)",
	Long: `A list of every password, showing the users, stored in the password manager.
You can use this list to see all the passwords you have stored or
to list every password for a specific service.

Examples:
	# List all the passwords stored
	passwordmanager list

	# List all the passwords stored for the service 'github'
	passwordmanager list -s github
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// I know that extracting the method of list passwords by service
		// and use it in the list all would sounds like an option, but that could
		// cause performance issues because each call to ListPasswordsByService
		// opens a new transaction and iterates through the usernames for that service.
		if service == "" {
			utils.ListAllPasswords()
		} else {
			utils.ListPasswordsByService(service)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&service, "service", "s", "", "Name of the service to list passwords for")
}
