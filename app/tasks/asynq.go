package tasks

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"github.com/hibiken/asynq"
	"os"
	"time"
)

//TaskAsynq -> TaskAsynq Struct
type TaskAsynq struct {
	Logger interfaces.Logger
	Server *asynq.Server
}

//NewTaskAsynq -> return new TaskAsynq struct,
func NewTaskAsynq(
	logger infrastracture.ArvanLogger,
) TaskAsynq {
	return TaskAsynq{
		Logger: &logger,
		Server: asynq.NewServer(asynq.RedisClientOpt{Addr: os.Getenv("RedisAddr")},
			asynq.Config{
				Concurrency: 10,
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"info":     1,
				},
			},
		),
	}
}

//NewClient -> return asynq client don't forget to close it
func (t *TaskAsynq) NewClient() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("RedisAddr")})
}

//NewScheduler -> return asynq periodic scheduler
func (t *TaskAsynq) NewScheduler() *asynq.Scheduler {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		panic(err)
	}
	return asynq.NewScheduler(
		asynq.RedisClientOpt{Addr: os.Getenv("RedisAddr")},
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)
}
