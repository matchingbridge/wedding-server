package match

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func AskMatch(c *gin.Context) {
	userID, partnerID := c.GetString("user_id"), c.Param("partner_id")

	match := models.Match{SenderID: userID, ReceiverID: partnerID}
	if err := mysql.DB.Create(&match).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("만남 요청 실패"))
		return
	}

	c.JSON(http.StatusOK, match)
}
