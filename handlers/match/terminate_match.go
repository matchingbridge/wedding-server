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

type TerminateMatchRequest struct {
	Message string `form:"message" json:"message" binding:"required"`
}

func TerminateMatch(c *gin.Context) {
	userID, partnerID := c.GetString("user_id"), c.Param("partner_id")

	var request TerminateMatchRequest
	if err := validateTerminateMatch(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	tx := mysql.DB.Begin()
	if err := tx.Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "데이터 베이스 연결 실패")
		return
	}
	defer tx.Rollback()

	var match models.Match
	query := "sender_id = ? and receiver_id = ?"
	if err := tx.Where(query, userID, partnerID).Or(query, partnerID, userID).Take(&match).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("매칭 정보를 찾을 수 없습니다"))
		return
	}

	if err := tx.Model(&match).Update("terminated_at", time.Now()).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("만남 종료 실패"))
		return
	}

	review := models.Review{
		MatchID:  match.MatchID,
		TargetID: partnerID,
		Content:  request.Message,
	}
	if err := tx.Create(&review).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("리뷰 데이터 생성 실패"))
		return
	}

	tx.Commit()

	c.Status(http.StatusOK)
}

func validateTerminateMatch(c *gin.Context, request *TerminateMatchRequest) error {
	if validationErr := c.ShouldBindJSON(&request); validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			if err.Field() == "message" {
				return errors.New("리뷰 내용을 다시 확인해주세요")
			}
		}
	}
	return nil
}
