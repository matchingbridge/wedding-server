package user

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetPartner(c *gin.Context) {
	partnerID := c.Param("partner_id")

	var partner models.User
	if err := mysql.DB.Where("user_id = ?", partnerID).Take(&partner).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("해당 유저 데이터를 가져올 수 없습니다"))
		return
	}
	c.JSON(http.StatusOK, partner)
}
