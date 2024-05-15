package taskrecover

import (
	"context"
	"time"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/botmain/types"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/task"
	pb "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/proto"
	tggateway "github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/tggateway/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type TaskRecoverer struct {
	DB           *ent.Client
	MessagesChan chan tggateway.TgMessageMsg
	ResponseChan chan tggateway.TgResponseMsg
	logger       *logrus.Logger
}

func NewTaskRecoverer(DB *ent.Client, MessagesChan chan tggateway.TgMessageMsg, ResponseChan chan tggateway.TgResponseMsg, logger *logrus.Logger) *TaskRecoverer {
	return &TaskRecoverer{
		DB:           DB,
		MessagesChan: MessagesChan,
		ResponseChan: ResponseChan,
		logger:       logger,
	}
}

func (t *TaskRecoverer) SetYieldTask(taskB *pb.TgYieldButton) error {
	if taskB.IsRestored {
		return nil
	}
	tasks, err := t.DB.Task.Query().Where(task.ID(taskB.GetUser().GetId())).All(context.Background())
	if err != nil {
		return err
	}
	if len(tasks) > 0 {
		data := pb.TooManyTasksResponse{
			User:  taskB.GetUser(),
			MsgId: taskB.GetMsgId(),
		}
		dataBytes, err := proto.Marshal(&data)
		if err != nil {
			return errors.Wrap(err, "failed marshal")
		}
		t.ResponseChan <- tggateway.TgResponseMsg{
			Name: "NotifyManyTasks",
			Data: dataBytes,
		}
		return &types.TooManyTasksError{Count: len(tasks)}
	}
	taskB.IsRestored = true
	data, err := proto.Marshal(taskB)
	if err != nil {
		return err
	}
	err = t.DB.Task.Create().
		SetName("Yield").
		SetData(data).
		SetID(taskB.GetUser().GetId()).OnConflict().DoNothing().Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRecoverer) ClearUserTasks(id int64) error {
	_, err := t.DB.Task.Delete().Where(task.ID(id)).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRecoverer) RecoverOnStartup() {
	tasks, err := t.DB.Task.Query().All(context.Background())

	if err != nil {
		t.logger.Error("failed get tasks", err)
	}
	for _, taskVal := range tasks {
		if taskVal.CreatedAt.Add(time.Hour * 2).Before(time.Now()) {
			err := t.DB.Task.DeleteOne(taskVal).Exec(context.Background())
			if err != nil {
				t.logger.Error("failed delete taskVal", err)
			}
			continue
		}
		switch taskVal.Name {
		case "Yield":
			//before 2 hours
			if time.Now().Before(taskVal.CreatedAt.Add(time.Hour * 2)) {
				t.MessagesChan <- tggateway.TgMessageMsg{
					Name: "Yield",
					Data: taskVal.Data,
				}
			}
		}
	}
}
