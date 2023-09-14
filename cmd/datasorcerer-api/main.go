package main

import (
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/r35krag0th/datasorcerer/internal"
	"github.com/r35krag0th/datasorcerer/internal/routehandlers"
	"go.uber.org/zap"
)

func main() {
	logger, err := middleware.NewDefaultLogger()
	if err != nil {
		panic(err)
	}

	logger.Debug("Creating Router")
	app := cli.NewRouter("DataSorcerer", "1.0.0")
	app.ServerLink("Local Dev", "http://127.0.0.1:8888")
	app.ServerLink("Production", "")

	db, err := internal.GetDatabaseConnection()
	if err != nil {
		logger.Fatal("failed to get database connection", zap.Error(err))
	}

	if err = internal.AutoMigrateModels(db); err != nil {
		logger.Fatal("failed to migrate", zap.Error(err))
	}

	rootResource := app.Resource("/")
	apiV1Resource := rootResource.SubResource("/api/v1")

	// hook up route handlers here
	routehandlers.RealmRoutes(apiV1Resource, db)

	logger.Debug("Listening")
	app.Run()
}
