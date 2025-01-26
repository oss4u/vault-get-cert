#!/usr/bin/env bash


docker compose up -d

export SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

export VAULT_TOKEN="root-token"
export VAULT_ADDR='https://127.0.0.1:8201'
export VAULT_SKIP_VERIFY='true'
mkdir ./tmp
export SHORT="int"
export PKI="pki-$SHORT"
export DOMAIN="int.sys-int.de"
export VAULT_DOMAIN=$VAULT_ADDR

vault secrets enable -path=$PKI pki

vault secrets tune -max-lease-ttl=8760h $PKI

vault write -field=certificate $PKI/root/generate/internal \
  common_name="sys-int.de" \
  issuer_name="root" \
  ttl=87600h > ./tmp/root_ca.crt

vault write $PKI/roles/int-sys-int\
    allow_localhost=true \
    allow_bare_domains=true \
    allowed_domains="localhost,sys-int.de,$DOMAIN" \
    allow_subdomains=true \
    max_ttl="720h"

vault write $PKI/config/cluster \
   path=$VAULT_DOMAIN/v1/$PKI \
   aia_path=$VAULT_DOMAIN/v1/$PKI

vault write $PKI/config/urls \
   issuing_certificates={{cluster_aia_path}}/issuer/{{issuer_id}}/der \
   crl_distribution_points={{cluster_aia_path}}/issuer/{{issuer_id}}/crl/der \
   ocsp_servers={{cluster_path}}/ocsp \
   enable_templating=true

vault secrets tune \
      -passthrough-request-headers=If-Modified-Since \
      -allowed-response-headers=Last-Modified \
      -allowed-response-headers=Location \
      -allowed-response-headers=Replay-Nonce \
      -allowed-response-headers=Link \
      $PKI

vault write $PKI/config/acme enabled=true

vault auth enable  approle

vault policy write server-default $SCRIPT_DIR/vault/policy_issue_certs.hcl

vault write auth/approle/role/server \
      policies="server-default" \
      token_ttl="1h" \
      token_max_ttl="4h"

# vault read auth/approle/role/server/role-id
# vault write -f auth/approle/role/server/secret-id
