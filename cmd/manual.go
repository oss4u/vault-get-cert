/*
Copyright © 2025 Marc Ende <me@e-beyond.de>
*/
package cmd

import (
	"fmt"
	"os"
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
	RunE: RunManualCmd,
}

func RunManualCmd(cmd *cobra.Command, args []string) error {
	err := configureSecrets(config)
	if err != nil {
		return fmt.Errorf("failed to configure secrets: %w", err)
	}
	err = internal.GetCertificates(config)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(manualCmd)

	flags := manualCmd.Flags()
	err := commonFlags(flags)
	if err != nil {
		fmt.Println("failed to bind flags")
		os.Exit(-1)
	}

}
