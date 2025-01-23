/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/stormy/utils"
	"github.com/spf13/cobra"
)

var setPathCmd = &cobra.Command{
	Use:   "set-path",
	Short: "Set the path for the database",
	Long: `Set the path for the database.
This command allows you to specify a new path and optionally force 
the change without confirmation.

This action will move the database to the new path if exists.

Example:
	# To set the path for the database to a new location
	stormy set-path
	# It will ask you for the new path and confirm the change`,

	Run: func(cmd *cobra.Command, args []string) {
		utils.SetDBPath(force)
	},
}

func init() {
	rootCmd.AddCommand(setPathCmd)

	setPathCmd.Flags().BoolVarP(&force, "force", "f", false, "Force the path change without confirmation")
}
