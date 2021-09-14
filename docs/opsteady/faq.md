# Frequently Asked Questions

* Why does Terraform give me an error while destroying the management infrastructure?

An administrator role assignment is created on the Key Vault when running the Terraform apply. When destroying the infrastructure, this role assignment is deleted first by Terraform and that will cause the remaining step (deleting the key vault) to fail. You can fix this by creating a role assignment manually via the portal, where you add yourself as administrator to the Key Vault. Afterwards you can run the Terraform destroy again and it will clean up the remaining infrastructure.
