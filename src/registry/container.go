package registry

import (
	"context"

	"github.com/dwdarm/go-url-shortener/src/configs"
	"github.com/dwdarm/go-url-shortener/src/handlers"
	"github.com/dwdarm/go-url-shortener/src/repositories"
	"github.com/dwdarm/go-url-shortener/src/services"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Container struct {
	ctn di.Container
}

func NewContainer(ctx context.Context) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	err = builder.Add((di.Def{
		Name: "global-setting",
		Build: func(ctn di.Container) (interface{}, error) {
			return configs.InitSetting(), nil
		},
	}))

	if err != nil {
		return nil, err
	}

	err = builder.Add(di.Def{
		Name: "db",
		Build: func(ctn di.Container) (interface{}, error) {
			setting := ctn.Get("global-setting").(*configs.Setting)

			if setting.UseMongodb {
				serverAPI := options.ServerAPI(options.ServerAPIVersion1)
				opts := options.Client().ApplyURI(setting.MongodbURI).SetServerAPIOptions(serverAPI)
				client, err := mongo.Connect(ctx, opts)
				if err != nil {
					panic(err)
				}

				db := client.Database(setting.MongodbName)

				return db, nil
			}

			return nil, nil
		},
	})

	if err != nil {
		return nil, err
	}

	err = builder.Add((di.Def{
		Name: "link-handler",
		Build: func(ctn di.Container) (interface{}, error) {
			setting := ctn.Get("global-setting").(*configs.Setting)

			var repo repositories.LinkRepository

			if setting.UseMongodb {
				repo = repositories.NewLinkMongodbRepository(
					ctn.Get("db").(*mongo.Database),
				)
			} else {
				repo = repositories.NewLinkMemoryRepository()
			}

			service := services.NewLinkService(repo)
			handler := handlers.NewLinkHandler(setting, service)

			return handler, nil
		},
	}))

	if err != nil {
		return nil, err
	}

	return &Container{ctn: builder.Build()}, nil

}

func (c *Container) GetLinkHandler() handlers.LinkHandler {
	return c.ctn.Get("link-handler").(handlers.LinkHandler)
}

func (c *Container) GetGlobalSetting() *configs.Setting {
	return c.ctn.Get("global-setting").(*configs.Setting)
}
