package config

import (
	"log"

	"github.com/spf13/viper"
)

/**
 * @brief      this function to get .env
 * @param      getenv  The getenv
 * @return     string
 */
func Env(getenv string) string {
	viper.SetConfigFile("./env.json")

	// Add paths. Accepts multiple paths. It will search these paths in given order
	viper.AddConfigPath("./")
	viper.AddConfigPath("$HOME/")

	// And then register config file name (no extension)
	viper.SetConfigName("env")

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	env := viper.Get(getenv)

	return env.(string)
}
