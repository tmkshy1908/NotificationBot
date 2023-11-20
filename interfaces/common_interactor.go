package interfaces

import (
	"context"
	"net/http"
)

type CommonInteractor interface {
	RootMain(context.Context, *http.Request) error
}
