package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/tmkshy1908/NotificationBot/interfaces"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

type ControllHandler struct {
	CommonController *interfaces.CommonController
}

func NewServer(lc line.LineClient) {
	c := &ControllHandler{
		CommonController: interfaces.NewController(lc),
	}
	NewRouter(c)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("##Server Connected##")
	}
}
