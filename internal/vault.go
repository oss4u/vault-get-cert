package internal

import (
	"context"
	"fmt"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"time"
)

func GetCertificates(config *Config) error {
	ctx := context.Background()
	client, err := authenticate(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}
	resp, err := client.Secrets.PkiIssuerIssueWithRole(
		ctx,
		config.PkiIssuer,
		config.PkiRole,
		schema.PkiIssuerIssueWithRoleRequest{
			CommonName:           config.ServerName,
			ExcludeCnFromSans:    false,
			IpSans:               config.IpAddresses,
			PrivateKeyFormat:     "",
			RemoveRootsFromChain: false,
			Ttl:                  config.CertTtl,
		},
		vault.WithMountPath(config.PkiPath),
	)
	if err != nil {
		return fmt.Errorf("failed to get certificate/key: %w", err)
	}
	err = WriteCertificate(config, resp.Data.Certificate)
	if err != nil {
		return fmt.Errorf("failed to write certificate: %w", err)
	}
	err = WriteCaChain(config, resp.Data.CaChain)
	if err != nil {
		return fmt.Errorf("failed to write ca-chain: %w", err)
	}
	err = WritePrivateKey(config, resp.Data.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to write private key: %w", err)
	}
	return nil
}

func authenticate(ctx context.Context, config *Config) (*vault.Client, error) {
	client, err := vault.New(
		vault.WithAddress(config.VaultAddress),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	resp, err := client.Auth.AppRoleLogin(
		ctx,
		schema.AppRoleLoginRequest{
			RoleId:   config.RoleID,
			SecretId: config.SecretID,
		},
		vault.WithMountPath(config.AppRolePath),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	if err := client.SetToken(resp.Auth.ClientToken); err != nil {
		return nil, fmt.Errorf("failed to set token: %w", err)
	}
	return client, nil
}
