package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

 type BaseGin gin.Engine


func ResponseWithJson(c *gin.Context, pyload interface{}){
	c.JSON(http.StatusOK, pyload)
}

func ResponseWithError(c *gin.Context, msg string){
	ResponseWithJson(c, map[string]string{"message": msg})
}
