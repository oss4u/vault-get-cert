/*
Copyright © 2025 Marc Ende <me@e-beyond.de>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"vault-get-cert/internal"
)

var (
	config  *internal.Config
	cfgFile string
)

func NewRootCommand() (*cobra.Command, error) {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
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
		RunE: RunCommand,
	}
	pflags := rootCmd.PersistentFlags()
	pflags.StringVarP(&cfgFile, "config", "c", "", "config file (default is /etc/vault-get-cert/config.yaml)")
	err := viper.BindPFlag("config", pflags.Lookup("config"))
	if err != nil {
		return nil, fmt.Errorf("failed to bind config flag: %w", err)
	}
	pflags.BoolVarP(&config.Debug, "debug", "d", false, "enable debugging")
	err = viper.BindPFlag("debug", pflags.Lookup("debug"))
	if err != nil {
		return nil, fmt.Errorf("failed to bind debug flag: %w", err)
	}

	pflags.BoolVarP(&config.SkipTlsVerify, "skip-tls-verify", "s", false, "skip tls verification")
	err = viper.BindPFlag("skip-tls-verify", pflags.Lookup("skip-tls-verify"))
	if err != nil {
		return nil, fmt.Errorf("failed to bind skip-tls-verify flag: %w", err)
	}

	flags := rootCmd.Flags()
	err = commonFlags(flags)
	if err != nil {
		return nil, fmt.Errorf("failed to add common flags: %w", err)
	}
	return rootCmd, nil
}

func RunCommand(cmd *cobra.Command, args []string) error {
	err := configureSecrets(config)
	if err != nil {
		return fmt.Errorf("failed to configure secrets: %w", err)
	}
	err = internal.RunServer(config)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {

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

	cobra.OnInitialize(initConfig)

	rootCmd, err := NewRootCommand()
	if err != nil {
		return fmt.Errorf("failed to create root command: %w", err)
	}
	manualCmd, err := NewManualCommand()
	if err != nil {
		return fmt.Errorf("failed to create manual command: %w", err)
	}
	initCmd, err := NewInitCommand()
	if err != nil {
		return fmt.Errorf("failed to create init command: %w", err)
	}
	rootCmd.AddCommand(manualCmd, initCmd)

	err = rootCmd.Execute()
	if err != nil {
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/vault-get-cert")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
