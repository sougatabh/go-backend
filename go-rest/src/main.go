package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Data struct {
	KEY   string `json:"key" binding:"required"`
	VALUE string `json:"value" binding:"required"`
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // no password set
		DB:       0,                                  // use default DB
	})

	router := gin.Default()
	router.POST("/put", func(c *gin.Context) {
		var data Data
		c.BindJSON(&data)
		err := rdb.Set(ctx, data.KEY, data.VALUE, 0).Err()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": "Data has been Pushed"})
	})

	router.GET("/get/:key", func(c *gin.Context) {
		var key = c.Param("key")
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)
		c.JSON(http.StatusOK, gin.H{"data": val})
	})

	router.Run()
}
