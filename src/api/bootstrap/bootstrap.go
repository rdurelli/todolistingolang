package bootstrap

import (
	"context"
	"github.com/rdurelli/todolist/src/api/controllers"
	"github.com/rdurelli/todolist/src/api/lib"
	"github.com/rdurelli/todolist/src/api/repositories"
	"github.com/rdurelli/todolist/src/api/routes"
	"github.com/rdurelli/todolist/src/api/services"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	lib.Module,
	controllers.Module,
	routes.Module,
	services.Module,
	repositories.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	env lib.Env,
	routes routes.Routes,
	logger lib.Logger,
	handler lib.RequestHandlerGin,
) {


	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")
			logger.Info("---------------------")
			logger.Info("------- CLEAN -------")
			logger.Info("---------------------")

			//conn.SetMaxOpenConns(10)
			go func() {
				//middlewares.Setup()
				routes.Setup()
				handler.Gin.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			//conn.Close()
			return nil
		},
	})
}