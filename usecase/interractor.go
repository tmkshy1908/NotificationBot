package usecase

import (
	"context"
	"net/http"
)

type CommonRepository interface {
	DivideEvent(context.Context, *http.Request) (string, string)
	CallReply(string, string)
	TimeSet(context.Context, string)
}
