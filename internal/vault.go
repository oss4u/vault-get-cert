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
	resp, err := client.Secrets.PkiIssuerIssueWithRole(
		ctx,
		"int-sys-int",
		"int-sys-int",
		schema.PkiIssuerIssueWithRoleRequest{
			AltNames:             "",
			CommonName:           config.ServerName,
			ExcludeCnFromSans:    false,
			Format:               "",
			IpSans:               nil,
			NotAfter:             "",
			OtherSans:            nil,
			PrivateKeyFormat:     "",
			RemoveRootsFromChain: false,
			SerialNumber:         "",
			Ttl:                  "",
			UriSans:              nil,
			UserIds:              nil,
		},
		vault.WithMountPath("pki-int"),
	)
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}
	fmt.Printf("Cert: %s", resp.Data.Certificate)
	fmt.Printf("Chain: %s", resp.Data.CaChain)
	fmt.Printf("Key: %s", resp.Data.PrivateKey)
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
