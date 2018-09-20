package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	"github.com/gin-contrib/cors"
	)

type NumberInfo struct {
	UserName string `json:"username"`
	Number   int    `json:"number"`
	IP       string `json:"ip"`
}

var InfoList []NumberInfo

func getNumberByRandom(ctx *gin.Context) {
	numberinfo := NumberInfo{}
	numberinfo.IP = ctx.ClientIP()
	username := ctx.DefaultQuery("username", "Anonymous")
	if username == "Anonymous"{
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	numberinfo.UserName = username
	numberinfo.Number = getRandomNumber()
	removeRepeat(numberinfo)
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"number":   numberinfo.Number,
	})
}

func removeRepeat(numberInfo NumberInfo) {
	for _, element := range InfoList {
		if element.IP == numberInfo.IP {
			return
		}
	}
	InfoList = append(InfoList, numberInfo)
}

func getStatistic(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, InfoList)
}

func clearStatistic(ctx *gin.Context) {
	InfoList = make([]NumberInfo, 0)
	ctx.JSON(http.StatusNoContent, nil)
}

func getRandomNumber() int {
	rand.Seed(int64(time.Now().UnixNano()))
	return rand.Intn(100) + 1
}

func main() {
	fmt.Println(getRandomNumber())
	router := gin.Default()
	router.GET("/api/v1alpha1/number", getNumberByRandom)
	router.GET("/api/v1alpha1/result", getStatistic)
	router.DELETE("/api/v1alpha1/result", clearStatistic)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type"}
	router.Use(cors.New(config))

	router.Run(":8080")
}
