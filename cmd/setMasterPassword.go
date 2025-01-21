/*
Copyright © 2025 Javier Santos javier.jsm21@gmail.com
*/
package cmd

import (
	"github.com/javsanmar5/stormy/utils"
	"github.com/spf13/cobra"
)

var salt string

var setMasterPasswordCmd = &cobra.Command{
	Use:   "set-master-password [password]",
	Short: "Set a master password to secure your passwords",
	Long: `Before saving your passwords, you must start setting a master password.
This password will be used to encrypt and decrypt your passwords securely.
It's important to remember this password as it will be required to access your passwords.

We encourage you to add a salt to your password to make it more secure. 
This should be as random as possible

Example:
	If you wanna set hello123 as your master password, you can run:
	stormy set-master-password hello123`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.SetMasterPasword(args[0], salt)
	},
}

func init() {
	rootCmd.AddCommand(setMasterPasswordCmd)

	setMasterPasswordCmd.Flags().StringVarP(&salt, "salt", "s", "", "Add a salt to your master password")
}
