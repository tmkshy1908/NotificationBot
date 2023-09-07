package infrastructure

import (
	"fmt"
	"net/http"
)

func NewServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("##Server Connected##")
	}
}
