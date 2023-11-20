package usecase

import (
	"context"
	"net/http"

	"github.com/tmkshy1908/NotificationBot/domain"
)

type CommonRepository interface {
	Add(context.Context, *domain.UserTime) error
	Delete(context.Context, *domain.UserTime) error
	DivideEvent(context.Context, *http.Request) (*domain.UserMsg, error)
	DivideMessage(context.Context, *domain.UserMsg) (*domain.UserTime, bool, error)
	CallReply(*domain.UserMsg) error
	Alarm(context.Context, *domain.UserTime) error
}
