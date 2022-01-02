package auth

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func Review(c *gin.Context) {
	authID, err := strconv.Atoi(c.Param("auth_id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("잘못된 id입니다"))
		return
	}

	if err := mysql.DB.Model(&models.Auth{AuthID: authID}).Update("reviewed_at", time.Now()).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("회원 리뷰 실패"))
		return
	}

	c.Status(http.StatusOK)
}
