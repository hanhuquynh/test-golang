package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
)

var d *xorm.Engine

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("load env err:", err)
	}
	// engine, err := ConnectDb()
	// if err != nil {
	// 	log.Fatal("connect db err:", err)
	// }
	// d = engine
}

func main() {
	// if err := CreateTableUser(); err != nil {
	// 	log.Fatal("create table user err:", err)
	// }
	r := gin.Default()
	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Trang chủ"})
	})
	// r.POST("/user", func(ctx *gin.Context) {
	// 	_, err := d.Insert(&User{
	// 		Name:    "User " + fmt.Sprint(rand.IntN(9999-1)+1),
	// 		Age:     rand.IntN(100-1) + 1,
	// 		Created: time.Now().Unix(),
	// 	})

	// 	if err != nil {
	// 		ctx.JSON(500, gin.H{"message": err.Error()})
	// 	}
	// 	ctx.JSON(200, gin.H{"message": "success"})
	// })
	r.Run(":" + os.Getenv("PORT"))
}
