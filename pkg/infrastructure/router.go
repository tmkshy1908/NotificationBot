package infrastructure

import "net/http"

func NewRouter(controller *ControllHandler) {
	http.HandleFunc("/hello", controller.CommonController.Sayhello)
}
