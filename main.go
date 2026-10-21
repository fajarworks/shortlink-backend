package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fajarworks/shortlink-backend/internal/config"
	"github.com/fajarworks/shortlink-backend/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           			Shortlink App
// @version         			1.0
// @description     			Aplication for eventhub API
// @host      					localhost:8080
// @BasePath  					/
// @securityDefinitions.apikey 	BearerAuth
// @in 							header
// @name 						Authorization
// @description 				Bearer token used as identiity for accessing backend
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	db := config.NewPsqlDb(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	pool, err := db.Conn()
	if err != nil {
		log.Println("can't connect DB:", err)
	}
	defer pool.Close()
	fmt.Println("berhasil connect db")

	rc := config.NewRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_NAME"), os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_PORT"))
	rdb := rc.Connect()
	defer rdb.Close()
	status := rdb.Ping(context.Background())

	res, err := status.Result()
	if err != nil {
		fmt.Println("gagal connect redis : ", err)
		return
	}
	fmt.Println(res)
	fmt.Println("berhasil connect redis")
	r := gin.Default()
	router.MainRouter(r, pool)
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))

}
