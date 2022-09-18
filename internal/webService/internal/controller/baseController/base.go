package baseController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseGin gin.Engine

func ResponseWithJson(c *gin.Context, pyload interface{}) {
	c.JSON(http.StatusOK, pyload)
}

func ResponseWithError(c *gin.Context, msg interface{}) {
	ResponseWithJson(c, map[string]interface{}{"message": msg})
}
