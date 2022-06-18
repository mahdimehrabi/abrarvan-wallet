package tasks

import (
	"challange/app/interfaces"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewTaskAsynq),
	fx.Provide(NewWalletTask),
)

type Task interface {
	HandlesToMux() error
}

type Tasks struct {
	logger     interfaces.Logger
	taskAsynq  TaskAsynq
	walletTask WalletTask
}

func NewTasks(
	logger interfaces.Logger,
	taskAsynq TaskAsynq,
	walletTask WalletTask) Tasks {
	return Tasks{
		logger:     logger,
		taskAsynq:  taskAsynq,
		walletTask: walletTask,
	}
}

func (t *Tasks) HandleTasks() error {
	serverMux := asynq.NewServeMux()
	serverMux.HandleFunc(
		TypeWalletCodesUpdate,
		t.walletTask.HandleWalletCodesUpdateTask,
	)
	return t.taskAsynq.Server.Run(serverMux)
}
