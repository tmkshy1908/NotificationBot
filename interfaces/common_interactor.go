package interfaces

import (
	"context"
	"net/http"
)

type CommonInteractor interface {
	DivideMessage(context.Context, *http.Request)
	TimeAlarm(context.Context)
}
