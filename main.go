package main

import (
	"database/sql"
	"os"
	"repo-pattern-w-trx-management/handler"
	repository "repo-pattern-w-trx-management/repo"
	"repo-pattern-w-trx-management/route"
	"repo-pattern-w-trx-management/usecase"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	// init logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// init db
	connStr := "host=localhost port=5432 user=postgres dbname=local password=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(db)
	uc := usecase.NewUsecase(repo)
	handler := handler.NewHandler(log, uc)

	route.Register(e, handler)

	if err = e.Start(":10001"); err != nil {
		log.Fatal(err)
	}
}
