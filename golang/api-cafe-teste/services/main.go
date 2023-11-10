package services

import (
	"database/sql"
	"time"
)

var db *sql.DB

const dbTimeOut = time.Second * 3

type Models struct {
	Coffe        Coffe
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
