package match

import (
	"errors"
	"net/http"
	"time"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AnswerMatchRequest struct {
	Accept bool `form:"accept" json:"accept" binding:"required"`
}

func AnswerMatch(c *gin.Context) {
	userID, partnerID := c.GetString("user_id"), c.Param("partner_id")

	var request AnswerMatchRequest
	if err := validateAnswerMatch(c, &request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tx := mysql.DB.Begin()
	if err := tx.Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "데이터 베이스 연결 실패")
		return
	}
	defer tx.Rollback()

	var match models.Match
	if err := tx.Where("sender_id = ? and receiver_id = ?", partnerID, userID).Take(&match).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("매칭 정보를 찾을 수 없습니다"))
		return
	}

	if request.Accept {
		if err := tx.Model(&match).Update("matched_at", time.Now()).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New("요청 수락 실패"))
			return
		}
	} else {
		if err := tx.Model(&match).Update("terminated_at", time.Now()).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New("요청 거절 실패"))
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, match)
}

func validateAnswerMatch(c *gin.Context, request *AnswerMatchRequest) error {
	if validationErr := c.ShouldBindJSON(&request); validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			if err.Field() == "accept" {
				return errors.New("수락 여부를 다시 확인해주세요")
			}
		}
	}
	return nil
}
