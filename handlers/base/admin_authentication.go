package base

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminAuthentication(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	if os.Getenv("PORT") == "" {
		c.Header("Access-Control-Allow-Origin", "*")
	} else {
		c.Header("Access-Control-Allow-Origin", "http://wedding-op-tool.s3-website.ap-northeast-2.amazonaws.com")
	}
}
