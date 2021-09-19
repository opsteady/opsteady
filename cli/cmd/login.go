package cmd

import (
	"github.com/opsteady/opsteady/cli/cache"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/spf13/cobra"
)

var (
	role string

	loginCmd = &cobra.Command{
		Use:     "login",
		Aliases: []string{"l"},
		Short:   "Login to Vault",
		Long:    `Login to Vault`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info().Str("role", role).Msg("Trying to login")
			tokenCache, err := cache.NewFileCache(globalConfig.CacheFile, &logger)
			if err != nil {
				logger.Fatal().Err(err).Msg("Could not initialize token cache")
			}

			if _, err := vault.NewVault(globalConfig.VaultAddress, role, globalConfig.VaultInsecure, tokenCache, &logger); err != nil {
				logger.Fatal().Err(err).Msg("Could not login to Vault")
			}
			logger.Debug().Msg("Login successful")
		},
	}
)

func initLogin() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&role, "role", "r", "admin", "Role to use: platform-admin, platform-developer, platform-viewer")
}
