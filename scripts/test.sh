#!/usr/bin/env bash
export SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
export VAULT_ADDR="https://127.0.0.1:8201"
export VAULT_SKIP_VERIFY='true'
export APP_ROLE_ID=$(VAULT_TOKEN="root-token" vault read auth/approle/role/server/role-id -format=json | jq -r .data.role_id)
export APP_SECRET_ID=$(VAULT_TOKEN="root-token" vault write -f auth/approle/role/server/secret-id -format=json | jq -r .data.secret_id)
echo "APP_ROLE_ID: $APP_ROLE_ID"
echo "APP_SECRET_ID: $APP_SECRET_ID"
export VAULT_TOKEN=$(vault write auth/approle/login role_id=$APP_ROLE_ID secret_id=$APP_SECRET_ID -format=json | jq -r .auth.client_token)
$SCRIPT_DIR/../vault-get-cert manual \
    --pki-path=pki-int \
    --pki-issuer=default \
    --pki-role=int-sys-int \
    --vault-addr=https://127.0.0.1:8201 \
    --server-name=test.int.sys-int.de \
    --role-id=$APP_ROLE_ID \
    --secret-id=$APP_SECRET_ID \
    --approle-path  approle \
    --ca-chain-path=$SCRIPT_DIR/tmp/ca-chain.crt \
    --cert-path=$SCRIPT_DIR/tmp/cert.crt \
    --key-path=$SCRIPT_DIR/tmp/key.key \
    --skip-tls-verify
