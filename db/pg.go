package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
)

type DbInitiator interface {
	InitSql(host, port, user, dbName, password string) (*sql.DB, error)
}

type dbInitiator struct {
	log *logrus.Logger
}

func NewDbinitiator(l *logrus.Logger) DbInitiator {
	return &dbInitiator{
		log: l,
	}
}

func (d *dbInitiator) InitSql(host, port, user, dbName, password string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
