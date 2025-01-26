package cmd

import (
	"fmt"
	"github.com/99designs/keyring"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"vault-get-cert/internal"
	"vault-get-cert/internal/networking"
)

func commonFlags(flags *pflag.FlagSet) error {
	flags.StringVar(&config.RoleID, "role-id", "", "the role-id to use for authentication")
	err := viper.BindPFlag("role-id", flags.Lookup("role-id"))
	if err != nil {
		return fmt.Errorf("failed to bind role-id flag: %w", err)
	}

	flags.StringVar(&config.SecretID, "secret-id", "", "the secret-id to use for authentication")
	err = viper.BindPFlag("secret-id", flags.Lookup("secret-id"))
	if err != nil {
		return fmt.Errorf("failed to bind secret-id flag: %w", err)
	}

	flags.StringVar(&config.VaultAddress, "vault-addr", "", "the address of the vault server")
	err = viper.BindPFlag("vault-addr", flags.Lookup("vault-addr"))
	if err != nil {
		return fmt.Errorf("failed to bind vault-addr flag: %w", err)
	}
	viper.SetDefault("vault-addr", "https://127.0.0.1:8200")

	flags.StringVar(&config.AppRolePath, "approle-path", "", "the approle path")
	err = viper.BindPFlag("approle-path", flags.Lookup("approle-path"))
	if err != nil {
		return fmt.Errorf("failed to bind approle-path flag: %w", err)
	}
	viper.SetDefault("approle-path", "approle")

	flags.StringVar(&config.CronExpression, "cron-expression", "", "the cron expression")
	err = viper.BindPFlag("cron-expression", flags.Lookup("cron-expression"))
	if err != nil {
		return fmt.Errorf("failed to bind cron-expression flag: %w", err)
	}

	flags.StringVar(&config.ServerName, "server-name", "", "the server name")
	err = viper.BindPFlag("server-name", flags.Lookup("server-name"))
	if err != nil {
		return fmt.Errorf("failed to bind server-name flag: %w", err)
	}

	flags.StringVar(&config.CertPath, "cert-path", "", "the path to the certificate")
	err = viper.BindPFlag("cert-path", flags.Lookup("cert-path"))
	if err != nil {
		return fmt.Errorf("failed to bind cert-path flag: %w", err)
	}

	flags.StringVar(&config.KeyPath, "key-path", "", "the path to the private key")
	err = viper.BindPFlag("key-path", flags.Lookup("key-path"))
	if err != nil {
		return fmt.Errorf("failed to bind key-path flag: %w", err)
	}

	flags.StringVar(&config.CaChainPath, "ca-chain-path", "", "the path to the ca chain")
	err = viper.BindPFlag("ca-chain-path", flags.Lookup("ca-chain-path"))
	if err != nil {
		return fmt.Errorf("failed to bind ca-chain-path flag: %w", err)
	}
	viper.SetDefault("ca-chain-path", "/etc/ssl/private/server-full.crt")

	flags.StringVar(&config.PkiPath, "pki-path", "", "the pki path")
	err = viper.BindPFlag("pki-path", flags.Lookup("pki-path"))
	if err != nil {
		return fmt.Errorf("failed to bind pki-path flag: %w", err)
	}
	viper.SetDefault("pki-path", "pki")

	flags.StringVar(&config.PkiRole, "pki-role", "", "the pki role in vault")
	err = viper.BindPFlag("pki-role", flags.Lookup("pki-role"))
	if err != nil {
		return fmt.Errorf("failed to bind pki-role flag: %w", err)
	}
	viper.SetDefault("pki-role", "server")

	flags.StringVar(&config.PkiIssuer, "pki-issuer", "", "the pki issuer")
	err = viper.BindPFlag("pki-issuer", flags.Lookup("pki-issuer"))
	if err != nil {
		return fmt.Errorf("failed to bind pki-issuer flag: %w", err)
	}

	flags.StringSliceVar(&config.IpAddresses, "ip-addresses", []string{}, "the ip addresses used in the certificate")
	err = viper.BindPFlag("ip-addresses", flags.Lookup("ip-addresses"))
	if err != nil {
		return fmt.Errorf("failed to bind ip-addresses flag: %w", err)
	}
	ipaddresses, err := networking.GetIpAddresses()
	if err != nil {
		fmt.Printf("failed to get ip addresses: %s", err.Error())
	}
	viper.SetDefault("ip-addresses", ipaddresses)

	flags.StringVar(&config.CertTtl, "cert-ttl", "", "the certificate's ttl")
	err = viper.BindPFlag("cert-ttl", flags.Lookup("cert-ttl"))
	if err != nil {
		return fmt.Errorf("failed to bind cert-ttl flag: %w", err)
	}
	viper.SetDefault("cert-ttl", "168h")
	return nil
}

func readSecretsFromKeystore() (string, string, error) {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: "vault-get-cert",
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to open keyring: %w", err)
	}
	roleId, err := ring.Get("role-id")
	if err != nil {
		return "", "", fmt.Errorf("failed to get role-id from keyring: %w", err)
	}
	secureId, err := ring.Get("secure-id")
	if err != nil {
		return "", "", fmt.Errorf("failed to get secure-id from keyring: %w", err)
	}
	return string(roleId.Data), string(secureId.Data), nil
}

func configureSecrets(config *internal.Config) error {
	if len(config.RoleID) == 0 && len(config.SecretID) == 0 {
		roleId, secretId, err := readSecretsFromKeystore()
		if err != nil {
			return fmt.Errorf("failed to read from keystore: %w", err)
		}
		config.RoleID = roleId
		config.SecretID = secretId
	}
	return nil
}
