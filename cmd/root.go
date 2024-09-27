package cmd

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
)

var (
	cfgFile string
	cfg     model.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fullstack-examination-2024",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".fullstack-examination-2024" (without extension).
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	cfg = model.Config{
		APIServer:     model.Server{Enable: true, Port: 8080},
		SwaggerServer: model.Server{Enable: false, Port: 1314},
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}
}
