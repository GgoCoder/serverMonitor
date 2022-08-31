package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginInfo := make(map[string]string)
	err := c.BindJSON(&loginInfo)
	fmt.Println(loginInfo)
	if err != nil {
		ResponseWithError(c, "error body")
	}
	fmt.Println(loginInfo["userName"] == "GgoCoder")
	fmt.Println(loginInfo["passWord"] == "123456")
	if loginInfo["userName"] == "GgoCoder" && loginInfo["passWord"] == "123456" {
		ResponseWithJson(c, "login successfully")
		return
	}
	ResponseWithError(c, "user or passwd error, please try again")
}
