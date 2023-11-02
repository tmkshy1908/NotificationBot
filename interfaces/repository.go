package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/tmkshy1908/NotificationBot/domain"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/db"
	"github.com/tmkshy1908/NotificationBot/pkg/infrastructure/line"
)

type CommonRepository struct {
	DB  db.SqlHandler
	Bot line.LineClient
}

const (
	INSERT_USERS string = "insert into users (id, hour, minute) values ($1,$2,$3)"
	DELETE_USERS string = "delete from users where id = $1"
)

func (r *CommonRepository) Add(ctx context.Context, users *domain.Users) (err error) {
	err = r.DB.ExecWithTx(func(*sql.Tx) error {
		_, err = r.DB.Exec(ctx, INSERT_USERS, users.Id, users.Hour, users.Minute)
		if err != nil {
			fmt.Println(err, "##Rep_Exec##")
		}
		return err
	})
	return
}

func (r *CommonRepository) Delete(ctx context.Context, users *domain.Users) (err error) {
	err = r.DB.ExecWithTx(func(*sql.Tx) error {
		_, err = r.DB.Exec(ctx, DELETE_USERS, users.Id)
		if err != nil {
			fmt.Println(err, "##Rep_Delete##")
		}
		return err
	})
	return
}

var timeMatcher = regexp.MustCompile("([0-9]+)時([0-9]+)分")

func (r *CommonRepository) DivideEvent(ctx context.Context, req *http.Request) (msg string, userId string) {
	msg, userId = r.Bot.CathEvents(ctx, req)
	return
}

func (r *CommonRepository) DivideMessage(ctx context.Context, userId string, msg string) (users *domain.Users, b bool) {
	m := timeMatcher.FindStringSubmatch(msg)
	if len(m) == 3 {
		hour, err := strconv.Atoi(m[1])

		// return "何時ですか？"
		minute, err := strconv.Atoi(m[2])
		if err != nil {
			// return "何分ですか？"
		}
		// err = h.store.Set(userId, hour, minute)
		if err != nil {
			// return "時間の設定に失敗しました"
		}
		users = &domain.Users{Id: userId, Hour: hour, Minute: minute}
		msg = fmt.Sprintf("%v時%v分ですね。わかりました。", hour, minute)
		r.Bot.MsgReply(msg, userId)

		return users, true
	} else {
		return nil, false
	}

}

func (r *CommonRepository) CallReply(msg string, userId string) {
	r.Bot.MsgReply(msg, userId)
}

func (r *CommonRepository) Alarm(ctx context.Context, userId string, users *domain.Users) {
	if users != nil {
		for {
			now := time.Now()
			timeValue := time.Date(now.Year(), now.Month(), now.Day(), users.Hour, users.Minute, 0, 0, now.Location())
			if now.After(timeValue) {
				timeValue = timeValue.Add(24 * time.Hour)
				msg := "時間です"
				r.Bot.MsgReply(msg, users.Id)
				r.Delete(ctx, users)
			}
			time.Sleep(timeValue.Sub(now))
		}
	}
}
