package chat

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MakeChatRequest struct {
	Message string `form:"message" json:"message" binding:"required"`
}

func MakeChat(c *gin.Context) {
	userID := c.GetString("user_id")
	chat := chat(c, userID, true)
	c.JSON(http.StatusOK, chat)
}

func ReplyChat(c *gin.Context) {
	userID := c.Param("user_id")
	chat := chat(c, userID, false)
	// if strings.HasPrefix(chat.Message, "{") && strings.HasSuffix(chat.Message, "}") {
	// 	partnerID := string(chat.Message[1 : len(chat.Message)-1])
	// 	match := models.Match{SenderID: userID, ReceiverID: partnerID}
	// 	if err := mysql.DB.Create(&match).Error; err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, "새 매치 데이터 생성 실패")
	// 		return
	// 	}
	// }
	c.JSON(http.StatusOK, chat)
}

func chat(c *gin.Context, userID string, isTalking bool) *models.Chat {
	var request MakeChatRequest
	if err := validateMakeChat(c, &request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil
	}

	chat := models.Chat{
		UserID:    userID,
		IsTalking: isTalking,
		Message:   request.Message,
	}

	if err := mysql.DB.Create(&chat).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("채팅 데이터 생성 실패"))
		return nil
	}

	return &chat
}

func validateMakeChat(c *gin.Context, request *MakeChatRequest) error {
	if validationErr := c.ShouldBindJSON(&request); validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			if err.Field() == "message" {
				return errors.New("메시지를 다시 확인해 주세요")
			}
		}
	}
	return nil
}
