package usecase

import "net/http"

type CommonRepository interface {
	DivideEvent(*http.Request) (string, string)
	CallReply(string, string)
}
