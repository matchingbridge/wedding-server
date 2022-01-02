package user

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.GetString("user_id")

	var user models.User
	if err := mysql.DB.Where("user_id = ?", userID).Take(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("내 데이터를 가져올 수 없습니다"))
		return
	}
	c.JSON(http.StatusOK, user)
}
