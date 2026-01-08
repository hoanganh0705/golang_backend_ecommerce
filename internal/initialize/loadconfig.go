package initialize

import (
	"GolangBackendEcommerce/global"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Implementation for loading configuration goes here
	viper := viper.New()

	// Check if config file is provided as command-line argument
	if len(os.Args) > 1 {
		viper.SetConfigFile(os.Args[1])
	} else {
		viper.AddConfigPath("./config/") // path to config
		viper.SetConfigName("local")     // ten file
		viper.SetConfigType("yaml")
	}

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	// read server configuration
	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("JWT Key:", viper.GetString("security.jwt.key"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
