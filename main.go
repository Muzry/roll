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
	removeRepeat(numberinfo)
	ctx.String(http.StatusOK, fmt.Sprintf("用户%sRoll的点数为%d。", username, numberinfo.Number))
}

func removeRepeat(numberInfo NumberInfo){
	for _, element := range InfoList{
		if element.IP == numberInfo.IP {
			return
		}
	}
	InfoList = append(InfoList, numberInfo)

}

func getStatistic(ctx * gin.Context){
	templateString := "%d, IP地址为%s的用户%sRoll的点数为%d。\n"
	resultString:= ""
	if len(InfoList) > 0 {
		for index, element := range InfoList {
			resultString += fmt.Sprintf(templateString,index + 1, element.IP, element.UserName, element.Number)
		}
		ctx.String(http.StatusOK, resultString)
	} else{
		ctx.String(http.StatusOK, "没有记录")
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