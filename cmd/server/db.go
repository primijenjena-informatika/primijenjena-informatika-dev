package server

import (
	"context"
	"log"
	"primijenjena-informatika-dev/ent"

	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabaseClient() *ent.Client {
	client, err := ent.Open("sqlite3", viper.GetString("db_path"))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
