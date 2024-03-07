package usecase

import (
	"context"
	"net/http"

	"github.com/tmkshy1908/NotificationBot/domain"
)

type CommonRepository interface {
	Add(context.Context, *domain.UserStates) error
	Delete(context.Context, *domain.UserStates) error
	DivideEvent(context.Context, *http.Request) (*domain.UserStates, error)
	DivideMessage(context.Context, *domain.UserStates) (*domain.UserStates, error)
	CallReply(*domain.UserStates) error
	Alarm(context.Context, <-chan *domain.UserStates) error
}
