package cmd

import (
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/opsteady/opsteady/cli/cache"
	"github.com/opsteady/opsteady/cli/component"
	"github.com/opsteady/opsteady/cli/components"
	"github.com/opsteady/opsteady/cli/credentials"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/spf13/cobra"
)

var (
	componentFlag       string
	azureIDFlag         string
	awsIDFlag           string
	dryRunFlag          bool
	platformVersionFlag string
)

func executeComponent(cmd *cobra.Command, executeComponent func(c component.Component)) {
	stopWhenAwsOrAzureIdNotSpecified(cmd)
	// Cleaning the TMP folder is very important because it stores rendered files
	// If not cleaned you might apply wrong settings or files to wrong environments
	ensureTmpFolderExistsAndIsEmpty()

	logger.Debug().Msg("Find the component")
	comp, ok := components.Components[componentFlag]
	if !ok {
		logger.Fatal().Str("component", componentFlag).Msg("Could not find component")
	}

	logger.Trace().Msg("Initialize dependencies for the default component implementation")
	vaultCache, configCache := initializeCacheDependency()

	vaultImpl, err := vault.NewVault(globalConfig.VaultAddress, "platform-admin", globalConfig.VaultToken, globalConfig.VaultInsecure, vaultCache, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not initialize Vault")
	}

	defaultComponent := component.DefaultComponent{
		Vault:                      vaultImpl,
		Credentials:                credentials.NewCredentials(vaultImpl, vaultCache, &logger),
		ComponentConfig:            component.NewComponentConfig(configCache, vaultImpl, &logger),
		GlobalConfig:               globalConfig,
		Logger:                     &logger,
		TerraformBackendConfigPath: "/tmp/tf-backend-config", // this is the default filesystem location for the Terraform backend configuration
		DefaultDependencies:        []string{"management-bootstrap"},
		ComponentName:              componentFlag,
		ComponentFolder:            calculateComponentFolder(comp),
		DryRun:                     dryRunFlag,
		AwsID:                      awsIDFlag,
		AzureID:                    azureIDFlag,
		PlatformVersion:            platformVersionFlag,
		Terraform:                  "terraform",
		Helm:                       "helm",
		Kubectl:                    "kubectl",
	}

	logger.Debug().Msg("Run the the component")
	comp.Initialize(defaultComponent)
	defer comp.(component.Component).Clean()
	executeComponent(comp.(component.Component))
}

func stopWhenAwsOrAzureIdNotSpecified(cmd *cobra.Command) {
	if cmd.Use == "deploy" || cmd.Use == "destroy" {
		if azureIDFlag == "" && awsIDFlag == "" {
			logger.Fatal().Msgf("You need to specify an 'azure-id' or 'aws-id' flag for %s command", cmd.Use)
		}
	}
}

func ensureTmpFolderExistsAndIsEmpty() {
	logger.Debug().Msg("Cleanup TMP dir")
	if err := os.RemoveAll(globalConfig.TmpFolder); err != nil {
		logger.Fatal().Err(err).Msg("could not clean up the TMP dir")
	}
	logger.Trace().Msg("Create TMP dir")
	if err := os.Mkdir(globalConfig.TmpFolder, 0700); err != nil {
		logger.Fatal().Err(err).Str("dir", globalConfig.TmpFolder).Msg("could not create temporary directory")
	}
}

func calculateComponentFolder(comp component.Initialize) string {
	fullPath := reflect.ValueOf(comp).Elem().Type().PkgPath()
	withoutGithub := strings.ReplaceAll(fullPath, "github.com/opsteady/opsteady/", "")
	return path.Dir(withoutGithub) // strip cicd from the path
}

func initializeCacheDependency() (cache.Cache, cache.Cache) {
	logger.Info().Msg("Initialize cache")
	if cacheAllFlag {
		cacheAll, err := cache.NewFileCache(globalConfig.CacheFile, &logger)
		if err != nil {
			logger.Fatal().Err(err)
		}
		return cacheAll, cacheAll
	} else if cacheFlag {
		fileCache, err := cache.NewFileCache(globalConfig.CacheFile, &logger)
		if err != nil {
			logger.Fatal().Err(err)
		}
		cacheConfig, err := cache.NewCache(&logger)
		if err != nil {
			logger.Fatal().Err(err)

		}
		return fileCache, cacheConfig
	}

	cache, err := cache.NewCache(&logger)
	if err != nil {
		logger.Fatal().Err(err)

	}
	return cache, cache
}

// initComponentFlags initializes the flags used by all component commands like build/test/deploy/etc...
func initComponentFlags(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	cmd.Flags().StringVarP(&componentFlag, "component", "c", "", "Name of the component")
	cmd.Flags().StringVarP(&azureIDFlag, "azure-id", "", "", "Azure subscription ID")
	cmd.Flags().StringVarP(&awsIDFlag, "aws-id", "", "", "AWS Account ID")
	if cmd.Use == "deploy" {
		cmd.Flags().BoolVarP(&dryRunFlag, "dry-run", "", false, "Dry run the deployment")
	}
	cmd.Flags().StringVarP(&platformVersionFlag, "platform-version", "", "v0", "Platform version")
	_ = cmd.MarkFlagRequired("component")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the component",
	Long:  `Deploy the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(cmd, func(c component.Component) {
			c.Deploy()
		})
	},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the component",
	Long:  `Destroy the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(cmd, func(c component.Component) {
			c.Destroy()
		})
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test the component",
	Long:  `Test the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(cmd, func(c component.Component) {
			c.Test()
		})
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the component",
	Long:  `Validate the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(cmd, func(c component.Component) {
			c.Validate()
		})
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the component",
	Long:  `Build the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(cmd, func(c component.Component) {
			c.Build()
		})
	},
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release the component",
	Long:  `Release the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(cmd, func(c component.Component) {
			c.Release()
		})
	},
}
