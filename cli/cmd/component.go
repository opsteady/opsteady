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
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
	"github.com/spf13/cobra"
)

const (
	readWriteExecute = 0700
)

var (
	componentFlag       string
	groupFlag           string
	azureIDFlag         string
	awsIDFlag           string
	localIDFlag         string
	dryRunFlag          bool
	platformVersionFlag string
	currentTarget       component.Target
	platformID          string
)

func executeComponent(executeComponent func(c component.Component)) {
	setTargetAndPlatformID()

	// Cleaning the TMP folder is very important because it stores rendered files
	// If not cleaned you might apply wrong settings or files to wrong environments
	ensureTmpFolderExistsAndIsEmpty()

	logger.Debug().Str("target", string(currentTarget)).Str("group", groupFlag).Str("component", componentFlag).Msg("Find the component")
	comp := components.Targets.FindComponent(currentTarget, component.Group(groupFlag), componentFlag)

	if comp == nil {
		logger.Fatal().Str("component", componentFlag).Msg("Could not find component")
	}

	logger.Trace().Msg("Initialize dependencies for the default component implementation")

	vaultCache, configCache := initializeCacheDependency()

	vaultImpl, err := vault.NewVault(globalConfig.VaultAddress, "platform-admin", globalConfig.VaultToken, globalConfig.VaultInsecure, vaultCache, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not initialize Vault")
	}

	meta := comp.GetMetadata()
	meta.AddRequiresInformationFrom(managementBootstrap.Instance.GetMetadata())

	defaultComponent := component.DefaultComponent{
		Vault:                      vaultImpl,
		Credentials:                credentials.NewCredentials(vaultImpl, vaultCache, &logger),
		ComponentConfig:            component.NewComponentConfig(configCache, vaultImpl, &logger),
		GlobalConfig:               globalConfig,
		Logger:                     &logger,
		TerraformBackendConfigPath: "/tmp/tf-backend-config", // this is the default filesystem location for the Terraform backend configuration
		Metadata:                   comp.GetMetadata(),
		ComponentFolder:            calculateComponentFolder(comp),
		CurrentTarget:              currentTarget,
		DryRun:                     dryRunFlag,
		PlatformID:                 platformID,
		PlatformVersion:            platformVersionFlag,
		Terraform:                  "terraform",
		CRD:                        "crd",
		KubeSetup:                  "kube_setup",
		Helm:                       "helm",
		KubePostSetup:              "kube_post_setup",
		Docker:                     "docker",
	}

	logger.Debug().Msg("Run the component")
	comp.Configure(defaultComponent)

	defer comp.Clean()
	executeComponent(comp)
}

func checkIDs(cmd *cobra.Command) { //nolint:cyclop
	if azureIDFlag == "" && awsIDFlag == "" && localIDFlag == "" {
		logger.Fatal().Msgf("You need to specify an 'azure-id' or 'aws-id' or 'local-id' flag for %s command", cmd.Use)
	}

	if azureIDFlag != "" && (awsIDFlag != "" || localIDFlag != "") {
		logger.Fatal().Msg("You can not mix 'azure-id' flag with 'aws-id' or 'local-id'")
	}

	if awsIDFlag != "" && (azureIDFlag != "" || localIDFlag != "") {
		logger.Fatal().Msg("You can not mix 'aws-id' flag with 'azure-id' or 'local-id'")
	}

	if localIDFlag != "" && (awsIDFlag != "" || azureIDFlag != "") {
		logger.Fatal().Msg("You can not mix 'local-id' flag with 'aws-id' or 'azure-id'")
	}
}

func setTargetAndPlatformID() {
	currentTarget = ""
	platformID = ""

	if localIDFlag != "" {
		currentTarget = component.TargetLocal
		platformID = localIDFlag
	}

	if awsIDFlag != "" {
		currentTarget = component.TargetAws
		platformID = awsIDFlag
	}

	if azureIDFlag != "" {
		currentTarget = component.TargetAzure
		platformID = azureIDFlag
	}

	if azureIDFlag == "management" {
		currentTarget = component.TargetManagement // Override this and set to management instead of azure
		platformID = azureIDFlag
	}
}

func ensureTmpFolderExistsAndIsEmpty() {
	logger.Debug().Msg("Cleanup TMP dir")

	if err := os.RemoveAll(globalConfig.TmpFolder); err != nil {
		logger.Fatal().Err(err).Msg("could not clean up the TMP dir")
	}

	logger.Trace().Msg("Create TMP dir")

	if err := os.Mkdir(globalConfig.TmpFolder, readWriteExecute); err != nil {
		logger.Fatal().Err(err).Str("dir", globalConfig.TmpFolder).Msg("could not create temporary directory")
	}
}

func calculateComponentFolder(comp component.Component) string {
	fullPath := reflect.ValueOf(comp).Elem().Type().PkgPath()
	withoutGithub := strings.ReplaceAll(fullPath, "github.com/opsteady/opsteady/", "")

	return path.Dir(withoutGithub) // strip cicd from the path
}

func initializeCacheDependency() (cache.Cache, cache.Cache) {
	logger.Info().Msg("Initialize cache")

	if cacheFlag {
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
	cmd.Flags().StringVarP(&groupFlag, "group", "g", "", "Name of the group the component is in")
	cmd.Flags().StringVarP(&azureIDFlag, "azure-id", "", "", "Azure platform ID")
	cmd.Flags().StringVarP(&awsIDFlag, "aws-id", "", "", "AWS platform ID")
	cmd.Flags().StringVarP(&localIDFlag, "local-id", "", "", "Local platform ID")

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
		executeComponent(func(c component.Component) {
			c.Deploy()
		})
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		checkIDs(cmd)
	},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the component",
	Long:  `Destroy the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(func(c component.Component) {
			c.Destroy()
		})
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		checkIDs(cmd)
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test the component",
	Long:  `Test the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(func(c component.Component) {
			c.Test()
		})
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the component",
	Long:  `Validate the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(func(c component.Component) {
			c.Validate()
		})
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the component",
	Long:  `Build the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(func(c component.Component) {
			c.Build()
		})
	},
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the component",
	Long:  `Publish the component`,
	Run: func(cmd *cobra.Command, args []string) {
		executeComponent(func(c component.Component) {
			c.Publish()
		})
	},
}
