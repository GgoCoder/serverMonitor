package LoginController

import (
	"fmt"
	base "serverMonitor/internal/webService/internal/controller/baseController"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginInfo := make(map[string]string)
	err := c.BindJSON(&loginInfo)
	fmt.Println(loginInfo)
	if err != nil {
		base.ResponseWithError(c, "error body")
	}
	fmt.Println(loginInfo["userName"] == "123456")
	fmt.Println(loginInfo["passWord"] == "123456")
	if loginInfo["userName"] == "123456" && loginInfo["passWord"] == "123456" {
		base.ResponseWithJson(c, "login successfully")
		return
	}
	base.ResponseWithError(c, "user or passwd error, please try again")
}
