package cmd

import (
	"context"

	"github.com/dwdarm/go-url-shortener/src/registry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	ctn, err := registry.NewContainer(context.TODO())
	if err != nil {
		panic(err)
	}

	setting := ctn.GetGlobalSetting()

	if setting.Env != "production" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))

	RegisterRoutes(ctn, r)

	return r
}
