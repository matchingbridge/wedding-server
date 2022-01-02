package match

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetMatch(c *gin.Context) {
	userID, partnerID := c.GetString("user_id"), c.Query("partner_id")

	var match models.Match
	if partnerID == "" {
		if err := mysql.DB.Where("sender_id = ? and terminated_at is null", userID).Or("receiver_id = ? and terminated_at is null", userID).Take(&match).Error; err != nil {
			c.AbortWithError(http.StatusNotFound, errors.New("매칭 정보를 찾을 수 없습니다"))
			return
		}
	} else {
		query := "sender_id = ? and receiver_id = ? and terminated_at is null"
		if err := mysql.DB.Where(query, userID, partnerID).Or(query, partnerID, userID).Take(&match).Error; err != nil {
			c.AbortWithError(http.StatusNotFound, errors.New("매칭 정보를 찾을 수 없습니다"))
			return
		}
	}

	c.JSON(http.StatusOK, match)
}
