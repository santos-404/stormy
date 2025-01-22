/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/stormy/utils"
	"github.com/spf13/cobra"
)

var force bool

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a password from the database",
	Long: `The "delete" command allows you to remove an existing password entry from the password database.

You need to specify the service and user associated with the password you want to delete. 
This command ensures that the corresponding entry is permanently removed from the database. 
Use this command with caution, as deleted passwords cannot be recovered.

Examples:
  # Delete the password of the user javier for the service github
  stormy delete -s github -u javier
  
  # Force delete without confirmation
  stormy delete -s github -u javier --force`,

	Run: func(cmd *cobra.Command, args []string) {
		utils.DeletePassword(service, username, force)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&service, "service", "s", "", "Name of the service (required)")
	deleteCmd.Flags().StringVarP(&username, "username", "u", "", "Username for the service (required)")
	deleteCmd.Flags().BoolVarP(&force, "force", "f", false, "Force the deletion without confirmation")

	deleteCmd.MarkFlagRequired("service")
	deleteCmd.MarkFlagRequired("username")
}
