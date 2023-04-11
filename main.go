package main

import (
	"backend-challenge/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=localhost user=root password=secret dbname=shop_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Db connection error", err)
	}

	fmt.Println("Database connected successfully")

	r := gin.Default()
	router.InitRoute(db, r)

	http.ListenAndServe(":8080", r)
}
