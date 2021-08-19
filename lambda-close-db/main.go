package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	runtime "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func handleRequest() (Response, error) {
	db, err := initDB()
	if err != nil {
		fmt.Println(err.Error())
		return Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	defer db.Close()

	title, err := getTitle(db)
	if err != nil {
		fmt.Println(err.Error())
		return Response{Message: err.Error(), StatusCode: http.StatusInternalServerError}, err
	}

	return Response{
		Message:    fmt.Sprintf("title: %s", *title),
		StatusCode: http.StatusOK,
	}, nil
}

func initDB() (*sql.DB, error) {
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	d, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", dbUser, dbPass, dbName))
	if err != nil {
		return nil, err
	}

	if err := d.Ping(); err != nil {
		return nil, err
	}

	return d, nil
}

func getTitle(db *sql.DB) (*string, error) {
	var title string
	sql := "SELECT title FROM sample_table LIMIT 1"
	err := db.QueryRow(sql).Scan(&title)
	if err != nil {
		return nil, err
	}

	return &title, nil
}

func main() {
	runtime.Start(handleRequest)
}
