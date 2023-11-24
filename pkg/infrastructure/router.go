package infrastructure

import "net/http"

func NewRouter(controller *ControllHandler) {
	http.HandleFunc("/hello", controller.CommonController.Sayhello)
	http.HandleFunc("/", controller.CommonController.LineHandller)
}
