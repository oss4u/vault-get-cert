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
	//resp, err := client.Write(ctx, "pki-int/issue/int-sys-int", map[string]interface{}{
	//	"common_name": config.ServerName,
	//	"ttl":         "48h",
	//	"format":      "pem",
	//})
	resp, err := client.Secrets.PkiWriteIssuer(ctx, "int-sys-int", schema.PkiWriteIssuerRequest{
		CrlDistributionPoints:        nil,
		EnableAiaUrlTemplating:       false,
		IssuerName:                   "",
		IssuingCertificates:          nil,
		LeafNotAfterBehavior:         "",
		ManualChain:                  nil,
		OcspServers:                  nil,
		RevocationSignatureAlgorithm: "",
		Usage:                        nil,
	}, vault.WithMountPath("pki-int"))
	//resp.Data.Certificate
	//resp.Data.CaChain
	//resp.Data.
	return nil
}

func authenticate(ctx context.Context, config *Config) (*vault.Client, error) {
	// prepare a client with the given base address
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
