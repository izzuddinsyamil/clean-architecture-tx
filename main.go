package main

import (
	"os"
	"repo-pattern-w-trx-management/db"
	"repo-pattern-w-trx-management/handler"
	repository "repo-pattern-w-trx-management/repo/pg"
	"repo-pattern-w-trx-management/route"
	"repo-pattern-w-trx-management/usecase"

	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// init db
	dbInit := db.NewDbinitiator(log)
	db, err := dbInit.InitSql(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		log.Fatal("Error connecting db", err)
	}

	repo := repository.NewRepo(db)
	uc := usecase.NewUsecase(repo)
	handler := handler.NewHandler(log, uc)

	route.Register(e, handler)

	if err = e.Start(":10001"); err != nil {
		log.Fatal(err)
	}
}
