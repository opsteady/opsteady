management_infra_acr_name                             = "devmgmtweu"
management_infra_vnet_address_space                   = ["10.0.0.0/19"]
management_infra_azure_subnet_pods_address_prefixes   = ["10.0.0.0/20"]
management_infra_platform_admins                      = []
management_infra_platform_admin_owners                = null
management_infra_platform_developers                  = []
management_infra_platform_developer_owners            = null
management_infra_platform_viewers                     = []
management_infra_platform_viewer_owners               = null
management_infra_domain                               = "os.opsteady.com"
management_infra_location                             = "westeurope"
management_infra_log_analytics_workspace_retention    = 30
management_infra_azure_subnet_public_address_prefixes = ["10.0.16.0/24"]

management_infra_aks_name                            = "management"
management_infra_aks_sku_tier                        = "Free"
management_infra_aks_kubernetes_version              = "1.21.2"
management_infra_aks_system_node_count               = 2
management_infra_aks_system_node_size                = "Standard_B2s"
management_infra_aks_api_server_authorized_ip_ranges = []
management_infra_aks_system_nodepool_node_count      = 3
management_infra_aks_system_nodepool_node_size       = "Standard_DS2_v2"

management_infra_key_vault_name           = "management"
management_infra_key_vault_ip_rules       = []
management_infra_key_vault_administrators = []
