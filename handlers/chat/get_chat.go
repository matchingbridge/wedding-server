package chat

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetChat(c *gin.Context) {
	userID := c.GetString("user_id")

	var chats []models.Chat
	if err := mysql.DB.Where("user_id = ?", userID).Find(&chats).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("채팅 데이터를 가져올 수 없습니다"))
		return
	}

	c.JSON(http.StatusOK, chats)
}
