package interfaces

import (
	"fmt"
	"net/http"

	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
	"github.com/tmkshy1908/NotificationBot/usecase"
)

type CommonController struct {
	Interactor CommonInteractor
}
type Controller interface {
	Sayhello(http.ResponseWriter, *http.Request)
}

func NewController(LineHandller line.LineClient) (cc *CommonController) {
	cc = &CommonController{
		Interactor: &usecase.CommonInteractor{
			CommonRepository: &CommonRepository{
				Bot: LineHandller,
			},
		},
	}
	return
}

func (cc *CommonController) Sayhello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ajaaaaaaa")
}
