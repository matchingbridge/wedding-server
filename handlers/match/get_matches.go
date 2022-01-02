package match

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	var matches []models.Match
	if err := mysql.DB.Find(&matches).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("매칭 데이터를 가져올 수 없습니다"))
		return
	}
	c.JSON(http.StatusOK, matches)
}
