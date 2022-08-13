package configs

import (
	"database/sql"
	"fmt"
	"hanoman-id/xendit-payment/internal/repository/query"
	"hanoman-id/xendit-payment/internal/v1/repositories"
	"hanoman-id/xendit-payment/internal/v1/usecase"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var oneUc sync.Once
var uc usecase.UseCase

func GetUsecase() usecase.UseCase {
	oneUc.Do(func() {
		uc = usecase.NewUseCase(
			getRepository(),
		)
	})

	return uc
}

var repo repositories.Repositories
var oneRepo sync.Once

func getRepository() repositories.Repositories {
	oneRepo.Do(func() {
		repo = repositories.NewRepositories(getQuery())
	})

	return repo
}

var qry *query.Queries
var oneSync sync.Once

func getQuery() *query.Queries {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oneSync.Do(func() {
		dbConnect := os.Getenv("DB_DSN")
		db, err := sql.Open("mysql", dbConnect)
		if err != nil {
			fmt.Println("Connector error: ", err)
			os.Exit(1)
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		qry = query.New(db)
	})

	return qry
}
