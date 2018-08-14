package main

import (
	"math/rand"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	)

type NumberInfo struct{
	UserName 	string
	Number		int
	IP			string
}

var InfoList []NumberInfo

func getNumberByRandom(ctx *gin.Context){
	numberinfo := NumberInfo{}
	numberinfo.IP = ctx.ClientIP()
	username := ctx.DefaultQuery("username", "Anonymous")
	numberinfo.UserName = username
	numberinfo.Number = getRandomNumber()
	InfoList = append(InfoList, numberinfo)
	ctx.String(http.StatusOK, fmt.Sprintf("UserName %s Roll %d", username, numberinfo.Number))
}

func getStatistic(ctx * gin.Context){
	templateString := "index %d, ip %s: %s roll %d\n"
	resultString:= ""
	if len(InfoList) > 0 {
		for index, element := range InfoList {
			resultString += fmt.Sprintf(templateString,index + 1, element.IP, element.UserName, element.Number)
		}
		ctx.String(http.StatusOK, resultString)
	} else{
		ctx.String(http.StatusOK, "nobody rolls.")
	}
}

func getRandomNumber() int {
	rand.Seed(int64(time.Now().UnixNano()))
	return rand.Intn(100) + 1
}


func main(){
	fmt.Println(getRandomNumber())
	router := gin.Default()
	router.GET("/api/v1alpha1/number", getNumberByRandom)
	router.GET("/api/v1alpha1/result", getStatistic)
	router.Run(":8080")
}