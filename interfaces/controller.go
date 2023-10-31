package interfaces

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/db"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
	"github.com/tmkshy1908/NotificationBot/usecase"
)

type CommonController struct {
	Interactor CommonInteractor
}
type Controller interface {
	Sayhello(http.ResponseWriter, *http.Request)
	LineHandller(http.ResponseWriter, *http.Request)
}

func NewController(SqlHandler db.SqlHandler, LineHandller line.LineClient) (cc *CommonController) {
	cc = &CommonController{
		Interactor: &usecase.CommonInteractor{
			CommonRepository: &CommonRepository{
				DB:  SqlHandler,
				Bot: LineHandller,
			},
		},
	}
	return
}

func (cc *CommonController) Sayhello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ajaaaaaaa")
}

func (cc *CommonController) LineHandller(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	cc.Interactor.RootMain(ctx, req)
}
