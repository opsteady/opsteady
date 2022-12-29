package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/opsteady/opsteady/cli/configuration"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "platform",
		Short: "Build/Deploy",
		Long:  `Use the cli to build and deploy components of Opsteady`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verboseFlag {
				logger = logger.Level(zerolog.DebugLevel)
			}
			if traceFlag {
				logger = logger.With().Caller().Stack().Logger().Level(zerolog.TraceLevel)
			}
		},
	}

	globalConfig *configuration.GlobalConfig
	logger       zerolog.Logger

	// Root flags (some are not configurable from config & ENV, maybe later)
	cacheFlag         bool
	verboseFlag       bool
	traceFlag         bool
	vaultFlag         string
	vaultInsecureFlag bool
)

func init() {
	initializeLogging()
	setDefaults()
	initializeGlobalFlags()
	initializeGlobalConfig()
	initLogin()
	initComponentFlags(buildCmd)
	initComponentFlags(deployCmd)
	initComponentFlags(destroyCmd)
	initComponentFlags(testCmd)
	initComponentFlags(validateCmd)
	initComponentFlags(publishCmd)
	initSetup()
}

func initializeLogging() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "15:04:05", // Only show the time not the date
	}
	logger = zerolog.New(output).With().Timestamp().Logger().Level(zerolog.InfoLevel)
}

func setDefaults() {
	homeDir, err := homedir.Dir()
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not find the home dir")
	}

	viper.SetDefault("cache_path", fmt.Sprintf("%s/.cache", homeDir))
	viper.SetDefault("cache_file", fmt.Sprintf("%s/.cache/.platform-cache", homeDir))
	viper.SetDefault("tmp_folder", "/tmp/opsteady")
}

func initializeGlobalFlags() {
	rootCmd.PersistentFlags().BoolVarP(&cacheFlag, "cache", "", cacheFlag, "Cache the passwords to reduce credential fetching overhead")
	rootCmd.PersistentFlags().StringVarP(&vaultFlag, "vault-address", "", vaultFlag, "Vault address")
	rootCmd.PersistentFlags().BoolVarP(&vaultInsecureFlag, "vault-insecure", "", vaultInsecureFlag, "Allow insecure Vault connection")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "", verboseFlag, "Verbose output")
	rootCmd.PersistentFlags().BoolVarP(&traceFlag, "trace", "", traceFlag, "Trace calls")

	// The following flags can be set using config file or ENV
	viperBindPFlag("vault_address", "vault-address")
	viperBindEnv("vault_address")
	viperBindPFlag("vault_insecure", "vault-insecure")
	viperBindEnv("vault_insecure")
	viperBindEnv("vault_token")
}

func viperBindPFlag(name, lookupName string) {
	if err := viper.BindPFlag(name, rootCmd.Flags().Lookup(lookupName)); err != nil {
		logger.Trace().Err(err).Msg("Failed to bind env with viper")
	}
}

func viperBindEnv(name string) {
	if err := viper.BindEnv(name); err != nil {
		logger.Fatal().Err(err).Msg("Failed to bind env with viper")
	}
}

func initializeGlobalConfig() {
	// Read the default config file
	viper.SetConfigFile("default-config.yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to read default config")
	}
	// Read the user specific config
	viper.SetConfigFile("config.yaml")

	if err := viper.MergeInConfig(); err != nil {
		if !strings.Contains(err.Error(), "no such file or directory") {
			logger.Fatal().Err(err).Msg("Failed to read user config")
		}

		logger.Warn().Msg("User config file not found, continuing without it")
	}

	// Use ENV
	viper.SetEnvPrefix("opsteady")
	viper.AutomaticEnv()

	globalConfig = &configuration.GlobalConfig{}

	if err := viper.Unmarshal(globalConfig); err != nil {
		logger.Fatal().Err(err).Msg("Failed to decode into GlobalConfig struct")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to execute command")
	}
}
