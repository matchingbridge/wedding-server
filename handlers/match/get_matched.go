package match

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetMatched(c *gin.Context) {
	userID := c.GetString("user_id")

	var matches []models.Match
	if err := mysql.DB.Where("sender_id = ? and matched_at is not null", userID).Or("receiver_id = ?", userID).Find(&matches).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("매칭 정보를 찾을 수 없습니다"))
		return
	}

	c.JSON(http.StatusOK, matches)
}
