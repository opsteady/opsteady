path "sys/leases/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

path "auth/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

path "sys/auth/*"
{
  capabilities = ["create", "update", "delete", "sudo"]
}

path "sys/auth"
{
  capabilities = ["read"]
}

path "sys/policies/acl"
{
  capabilities = ["read","list"]
}

path "sys/policies/acl/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

path "secrets/data/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

path "secrets/metadatadata/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

path "sys/mounts/*"
{
  capabilities = ["create", "read", "update", "delete", "list", "sudo"]
}

path "sys/mounts"
{
  capabilities = ["read"]
}

path "sys/health"
{
  capabilities = ["read", "sudo"]
}

path "aws/*"
{
  capabilities = ["read", "create", "update", "delete", "list", "sudo"]
}

path "azure/*"
{
  capabilities = ["read", "create", "update", "delete", "list", "sudo"]
}

path "config/*"
{
  capabilities = ["read", "create", "update", "delete", "list", "sudo"]
}
