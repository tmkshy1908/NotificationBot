package main

import (
	"fmt"

	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/db"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

func main() {
	bot, err := line.NewLineClient()
	if err != nil {
		fmt.Println(err)
	}
	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}
	infrastructure.NewServer(db, bot)
}
