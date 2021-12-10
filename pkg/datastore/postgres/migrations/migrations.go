package main

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres"
)

// gorm it will be using auto migration and never delete any column and any data
// TODO: Improve Migrations strategies rolling up and rolling back
func main() {
	config.New("./secret.yaml")
	postgresClient := postgres.NewClient()

	postgresClient.Migrate()
}
