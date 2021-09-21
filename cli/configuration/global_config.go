package configuration

// GlobalConfig contains all the global config data for Opsteady
// Note: these are all set in the root of the cmd package
type GlobalConfig struct {
	VaultAddress             string `mapstructure:"vault_address"`
	VaultInsecure            bool   `mapstructure:"vault_insecure"`
	CachePath                string `mapstructure:"cache_path"`
	CacheFile                string `mapstructure:"cache_file"`
	TmpFolder                string `mapstructure:"tmp_folder"`
	TenantID                 string `mapstructure:"tenant_id"`
	ManagementSubscriptionID string `mapstructure:"management_subscription_id"`
}
