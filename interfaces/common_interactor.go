package interfaces

import "net/http"

type CommonInteractor interface {
	DivideMessage(*http.Request)
}
