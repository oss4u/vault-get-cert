/*
Copyright © 2025 Marc Ende <me@e-beyond.de>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: RunInitCmd,
}

func RunInitCmd(cmd *cobra.Command, args []string) error {
	fmt.Println("init called")
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
	flags := initCmd.Flags()
	flags.StringVar(&config.RoleID, "role-id", "", "the role-id to use for authentication")
	err := viper.BindPFlag("role-id", flags.Lookup("role-id"))
	if err != nil {
		fmt.Println("failed to bind role-id flag")
		os.Exit(-1)
	}
	flags.StringVar(&config.SecretID, "secret-id", "", "the secret-id to use for authentication")
	err = viper.BindPFlag("secret-id", flags.Lookup("secret-id"))
	if err != nil {
		fmt.Println("failed to bind secret-id flag")
		os.Exit(-1)
	}
}
