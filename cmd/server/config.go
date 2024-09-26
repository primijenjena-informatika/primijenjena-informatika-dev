package server

import (
	"log"

	"github.com/spf13/viper"
)

func SetupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	// viper.AddConfigPath("/etc/codespaces/")
	viper.AddConfigPath(".")

	viper.SetDefault("port", 8080)
	viper.SetDefault("enable_tls", false)
	viper.SetDefault("db_path", "file:dev.db?_fk=1&_journal_mode=WAL")
	viper.SetDefault("github_codespace_template", "primijenjena-informatika/codespace-template")
	viper.SetDefault("github_org_token", "GITHUB_ORG_SECRET_TOKEN")
	viper.SetDefault("github_org_name", "GITHUB_ORG_NAME")
	viper.SetDefault("github_client_id", "GITHUB_OAUTH_APP_CLIENT_ID")
	viper.SetDefault("github_client_secret", "GITHUB_OAUTH_APP_CLIENT_SECRET")
	viper.SetDefault("codespace_retention_period", 30*24*60)
	viper.SetDefault("codespace_machine", "standardLinux")
	viper.SetDefault("codespace_branch", "develop")
	viper.SetDefault("redirect_uri", "http://localhost:8080/new")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("Error reading config file: %s", err)

			return
		}

		// Config file not found; create a default one
		err = viper.SafeWriteConfig()
		if err != nil {
			log.Fatalf("Error creating default config file: %s", err)
		}
		log.Println("Created default config file")
	}
}
