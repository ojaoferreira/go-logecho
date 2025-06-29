package repository

import (
	"go-logecho/model"

	"go.uber.org/zap"
)

type MessageRepository interface {
	CreateMessage(m model.Message)
}

type messageRepo struct {
	zap *zap.Logger
}

func NewMessageRepository(zap *zap.Logger) MessageRepository {
	return &messageRepo{zap}
}

func (z *messageRepo) CreateMessage(m model.Message) {
	z.zap.Info(m.Message)
}
