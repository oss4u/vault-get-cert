#!/usr/bin/env bash

export VAULT_ADDR="https://127.0.0.1:8201"
export VAULT_SKIP_VERIFY='true'
export APP_ROLE_ID=$(VAULT_TOKEN="root-token" vault read auth/approle/role/server/role-id -format=json | jq -r .data.role_id)
export APP_SECRET_ID=$(VAULT_TOKEN="root-token" vault write -f auth/approle/role/server/secret-id -format=json | jq -r .data.secret_id)
echo "APP_ROLE_ID: $APP_ROLE_ID"
echo "APP_SECRET_ID: $APP_SECRET_ID"
export VAULT_TOKEN=$(vault write auth/approle/login role_id=$APP_ROLE_ID secret_id=$APP_SECRET_ID -format=json | jq -r .auth.client_token)
vault write pki-int/issue/int-sys-int common_name="test.sys-int.de" ttl="1h"
