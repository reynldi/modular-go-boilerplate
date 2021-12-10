package application

import (
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/application/router"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/repository"
	"github.com/reynldi/modular-go-boilerplate/pkg/delivery/rest"
	"github.com/reynldi/modular-go-boilerplate/pkg/helper"
	"gorm.io/gorm"
)

type AppClient struct {
	Db *gorm.DB
}

func LoadApp(r *gin.Engine) AppClient {
	pg := postgres.NewClient()
	pg.Migrate()
	repo := repository.NewRepository(pg.Db)

	globalHandler := rest.NewHandler()
	router.SetupRouter(r, globalHandler)

	helper.InitUptime()

	// Register the Module
	RegisterModule(regModuleOptions{
		r,
		repo,
	})

	return AppClient{
		Db: pg.Db,
	}
}

func StopApp() {
	// TODO: graceful shutdown from SIGINT
}
