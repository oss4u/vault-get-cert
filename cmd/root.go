/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/99designs/keyring"
	"github.com/Mrucznik/wonsz"
	"github.com/spf13/cobra"
	"os"
	"vault-get-cert/internal"
)

var config *internal.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vault-get-cert",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceUsage: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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
		err = internal.RunServer(config)
		if err != nil {
			return fmt.Errorf("failed to run server: %w", err)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config = &internal.Config{
		RoleID:         "",
		SecretID:       "",
		VaultAddress:   "",
		AppRolePath:    "approle",
		CronExpression: "",
		ServerName:     "",
		CertPath:       "/etc/ssl/private/server.crt",
		KeyPath:        "/etc/ssl/private/server.key",
		CaChainPath:    "/etc/ssl/private/server-full.crt",
	}
	err := wonsz.BindConfig(config, rootCmd,
		wonsz.ConfigOpts{
			ConfigPaths: []string{".", "/etc/vault-get-cert"},
			ConfigType:  "yaml",
			ConfigName:  "config",
		},
	)
	if err != nil {
		return
	}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vault-get-cert.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
