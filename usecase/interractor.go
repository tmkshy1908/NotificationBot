package usecase

import (
	"context"
	"net/http"

	"github.com/tmkshy1908/NotificationBot/domain"
)

type CommonRepository interface {
	Add(context.Context, *domain.Users) error
	Delete(context.Context, *domain.Users) error
	DivideEvent(context.Context, *http.Request) (string, string)
	DivideMessage(context.Context, string, string) (*domain.Users, bool)
	CallReply(string, string)
	// TimeSet(string, string)
	Alarm(context.Context, string, *domain.Users)
}
