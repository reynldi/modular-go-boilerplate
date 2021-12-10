package main

import (
	"flag"
	"fmt"
	"github.com/reynldi/modular-go-boilerplate/cmd/graphql"

	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/application"
	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
)


func main() {
	loadEnvConfig()
	
	r := gin.Default()

	// load application instance
	application.LoadApp(r)

	// Setup Graphql Server
	r.POST("/graph/query", graphql.GqlServer())
	r.GET("/graph", graphql.GqlPlayground())

	// running http server on `port`
	r.Run(fmt.Sprintf(`:%d`, config.GetConfig().Application.Port))
}

func loadEnvConfig(){
	var configFile string
	flag.StringVar(&configFile, "config", "./secret.yaml", "path to config file")
	flag.Parse()

	if err := config.ValidateConfigPath(configFile); err != nil {
		fmt.Println(err.Error())
	}

	_, err := config.New(configFile)
	if err != nil {
		fmt.Println(err.Error())
	}
}