package postgres

import (
	"fmt"
	"log"

	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/entity"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ClientInstance struct {
	Db *gorm.DB
}

func NewClient() *ClientInstance {
	dbConfig := config.GetConfig().Postgres
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta`, dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("Database Error:" + err.Error())
		return nil
	}

	log.Println("Postgres database connected")

	return &ClientInstance{Db: db}
}

func (client *ClientInstance) Migrate() {
	client.Db.AutoMigrate(
		&entity.AuthUser{},
	)
	log.Println("Database migrated")
}
