# Frequently Asked Questions

* Why does Terraform give me an authorization error on the Key Vault while destroying the management infrastructure?

An administrator role assignment is created on the Key Vault when running the Terraform apply. When destroying the infrastructure, this role assignment is deleted first by Terraform and that will cause the remaining step (deleting the key vault) to fail. You can fix this by creating a role assignment manually via the portal, where you add yourself as administrator to the Key Vault. Afterwards you can run the Terraform destroy again and it will clean up the remaining infrastructure.

* Why does Terraform give me an error on connecting to the Kubernetes cluster while destroying the management infrastructure?

We are configuring the Kubernetes and Helm providers with authentication attributes from the AKS cluster resource. When performing a destroy these attributes are not available anymore due the way Terraform works internally. The workaround is to not refresh the resource when performing the destroy (add `-refresh=false` to the destroy command)
