package base

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func UserAuthentication(c *gin.Context) {
	authID, err := getReqIdFromToken(c.Request)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, errors.New("잘못된 접근입니다"))
		return
	}

	var auth models.Auth
	if err := mysql.DB.Where("auth_id = ?", authID).Take(&auth).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("가입 정보를 찾을 수 없습니다"))
		return
	}

	c.Set("user_id", auth.UserID)
	c.Next()
}
