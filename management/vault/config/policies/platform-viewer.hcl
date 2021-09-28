path "sys/leases/*"
{
  capabilities = ["read", "list"]
}

path "auth/*"
{
  capabilities = ["read", "list"]
}

path "sys/auth"
{
  capabilities = ["read"]
}

path "sys/policies/acl"
{
  capabilities = ["read"]
}

path "sys/policies/acl/*"
{
  capabilities = ["read", "list"]
}

path "secret/data/*"
{
  capabilities = ["read"]
}

path "secret/metadata/*"
{
  capabilities = ["list"]
}

path "sys/mounts/*"
{
  capabilities = ["read", "list"]
}

path "sys/mounts"
{
  capabilities = ["read"]
}

path "sys/health"
{
  capabilities = ["read", "sudo"]
}
