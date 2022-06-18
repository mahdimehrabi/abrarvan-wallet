package app

import (
	"challange/app/controller"
	"challange/app/infrastracture"
	"challange/app/repository"
	"challange/app/routes"
	"challange/app/services"
	"challange/app/tasks"
	"context"
	"go.uber.org/fx"
	"log"
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
	tasks.Module,
	fx.Invoke(Bootstrap),
)

func Bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastracture.ArvanLogger,
	db infrastracture.PgxDB,
	taskAsynq tasks.TaskAsynq,
	walletRoutes routes.WalletRoutes,
	walletTask tasks.WalletTask,
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

	//asynq task handler
	go func() {
		tasksStruct := tasks.NewTasks(&logger, taskAsynq, walletTask)
		err := tasksStruct.HandleTasks()
		if err != nil {
			logger.Error("Failed to run asynq handlers:" + err.Error())
		}
	}()
	//periodic task scheduler
	go func() {
		task, err := walletTask.NewUpdateCodesWalletTask()
		if err != nil {
			logger.Error("Failed to start task for counting segments:" + err.Error())
		}
		scheduler := taskAsynq.NewScheduler()
		entryID, err := scheduler.Register("@every 30s", task)
		if err != nil {
			logger.Error("Failed to start task for counting segments:" + err.Error())
		}
		log.Printf("registered an entry: %q\n", entryID)

		if err := scheduler.Run(); err != nil {
			logger.Error("Failed to start task for counting segments:" + err.Error())
		}
	}()
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
