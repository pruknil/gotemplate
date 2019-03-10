package httpsvr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func listUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":   "",
		"payload": ""})

}
