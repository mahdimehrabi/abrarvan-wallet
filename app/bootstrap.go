package app

import (
	"challange/app/controller"
	"challange/app/infrastracture"
	"challange/app/repository"
	"challange/app/routes"
	"challange/app/services"
	"context"
	"go.uber.org/fx"
	"net/http"
	"os"
	"time"
)

var BootstrapModule = fx.Options(
	infrastracture.Module,
	repository.Module,
	services.Module,
	controller.Module,
	routes.Module,
	fx.Invoke(Bootstrap),
)

func Bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastracture.ArvanLogger,
	db infrastracture.PgxDB,
	walletRoutes routes.WalletRoutes,
) {
	port := os.Getenv("ServerPort")

	// create a new serve mux and register routes
	sm := http.NewServeMux()

	walletRoutes.AddRoutes(sm)
	// create a new server
	s := http.Server{
		Addr:         ":" + port,        // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     logger.LG,         // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			// start the server
			go func() {
				logger.LG.Println("Starting server on port " + port)

				err := s.ListenAndServe()
				if err != nil {
					logger.LG.Printf("Error starting server: %s\n", err)
					os.Exit(1)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})

}
