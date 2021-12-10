package main

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/seeds"
	"log"

	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres"
)

func main() {
	config.New("./secret.yaml")
	postgresClient := postgres.NewClient()

	users := seeds.SeedAuthUsers(postgresClient.Db).CreateMany(10)

	log.Printf("%d Auth user seeds", len(users))
}
