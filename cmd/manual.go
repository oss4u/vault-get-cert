/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/99designs/keyring"
	"vault-get-cert/internal"

	"github.com/spf13/cobra"
)

// manualCmd represents the manual command
var manualCmd = &cobra.Command{
	Use:   "manual",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ring, err := keyring.Open(keyring.Config{
			ServiceName: "vault-get-cert",
		})
		if err != nil {
			return fmt.Errorf("failed to open keyring: %w", err)
		}
		roleId, err := ring.Get("role-id")
		if err != nil {
			return fmt.Errorf("failed to get role-id from keyring: %w", err)
		}
		secureId, err := ring.Get("secure-id")
		if err != nil {
			return fmt.Errorf("failed to get secure-id from keyring: %w", err)
		}
		config.RoleID = string(roleId.Data)
		config.SecretID = string(secureId.Data)
		err = internal.GetCertificates(config)
		if err != nil {
			return fmt.Errorf("failed to run server: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(manualCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manualCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manualCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
