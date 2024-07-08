package service

import "context"

var _ IMessageService = (*MessageService)(nil)

type IMessageService interface {
	PushLoginMessage(ctx context.Context, id uint)
}

type MessageService struct {
}

func (m *MessageService) PushLoginMessage(ctx context.Context, id uint) {

}
