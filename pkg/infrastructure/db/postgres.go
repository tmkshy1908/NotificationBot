package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type SqlConf struct {
	Conn *sql.DB
}

type dbSettings struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

type SqlHandler interface {
	Exec(context.Context, string, ...interface{}) (sql.Result, error)
	Query(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(context.Context, string, ...interface{}) *sql.Row
	ExecWithTx(txFunc func(*sql.Tx) error) error
}

func NewHandler() (h SqlHandler, err error) {
	conf := dbSettings{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", conf.Host, conf.User, conf.Password, conf.Port, conf.Name)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		err = db.Ping()
		if err != nil {
			fmt.Println("DB接続エラー: ", err)
		}
	}
	fmt.Println("DB Connected.")
	h = &SqlConf{Conn: db}

	return
}

func (h *SqlConf) Exec(ctx context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	res, err = h.Conn.ExecContext(ctx, query, args...)
	return
}

func (h *SqlConf) Query(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = h.Conn.QueryContext(ctx, query, args...)
	return
}

func (h *SqlConf) QueryRow(ctx context.Context, query string, args ...interface{}) (row *sql.Row) {
	row = h.Conn.QueryRowContext(ctx, query, args...)
	return
}

func (h *SqlConf) ExecWithTx(txFunc func(*sql.Tx) error) (err error) {
	tx, err := h.Conn.Begin()
	if err != nil {
		log.Println("##withTx##")
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(err, "##Rollback##")
			err = tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println(err, "##Rollback err##")
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = rollbackErr
			}
		} else {
			fmt.Println("##commit##")
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return
}
