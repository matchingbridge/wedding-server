package user

import (
	"net/http"
	"wedding/datastore/s3"

	"github.com/gin-gonic/gin"
)

func GetMedia(c *gin.Context) {
	bucket := c.Param("bucket")
	userID := c.Param("user_id")

	url, err := s3.DownloadFile(s3.MakeKey(bucket, userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "파일을 가져올 수 없습니다.")
	}

	c.Redirect(http.StatusTemporaryRedirect, url)
}
