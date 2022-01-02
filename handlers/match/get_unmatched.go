package match

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetUnmatched(c *gin.Context) {
	userID := c.GetString("user_id")

	var matches []models.Match
	query := mysql.DB.Where("receiver_id = ? and asked_at is not null and matched_at is null", userID)
	if err := query.Find(&matches).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("매칭 정보를 찾을 수 없습니다"))
		return
	}

	c.JSON(http.StatusOK, matches)
}
