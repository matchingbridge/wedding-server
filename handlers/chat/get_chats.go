package chat

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetChats(c *gin.Context) {
	var chats []models.Chat
	if err := mysql.DB.Find(&chats).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("채팅 데이터를 가져올 수 없습니다"))
		return
	}
	c.JSON(http.StatusOK, chats)
}
