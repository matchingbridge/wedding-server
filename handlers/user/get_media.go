package user

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/datastore/s3"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetMedia(c *gin.Context) {
	bucket := c.Param("bucket")
	userID := c.Param("user_id")

	var user models.User
	if err := mysql.DB.Where("user_id = ?", userID).Take(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("상대방 데이터를 가져올 수 없습니다"))
		return
	}

	var image string
	if bucket == "face1" {
		image = user.Face1
	} else if bucket == "face2" {
		image = user.Face2
	} else if bucket == "body1" {
		image = user.Body1
	} else if bucket == "body2" {
		image = user.Body2
	} else if bucket == "video" {
		image = user.Video
	} else {
		c.AbortWithError(http.StatusBadRequest, errors.New("알 수 없는 이미지 형식"))
		return
	}

	url, err := s3.DownloadFile(s3.MakeKey(bucket, image))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "파일을 가져올 수 없습니다.")
	}

	c.Redirect(http.StatusTemporaryRedirect, url)
}
