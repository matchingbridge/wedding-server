package suggestion

import (
	"errors"
	"net/http"
	"time"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSuggestions(c *gin.Context) {
	userID := c.GetString("user_id")

	tx := mysql.DB.Begin()
	if err := tx.Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "데이터 베이스 연결 실패")
		return
	}
	defer tx.Rollback()

	var suggestions []models.Suggestion

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	if err := tx.Where("user_id = ? and year(suggested_at) = ? and month(suggested_at) = ? and day(suggested_at) = ?", userID, year, month, day).Take(&suggestions).Error; err != nil {
		if findErr := find(userID, tx); findErr != nil {
			c.JSON(http.StatusOK, suggestions)
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, suggestions)
}

func find(userID string, tx *gorm.DB) error {
	var user models.User
	if searchErr := tx.Select("users.gender, users.area").Where("user_id = ?", userID).First(&user).Error; searchErr != nil {
		return errors.New("유저 정보 찾기 실패")
	}

	var users []models.User
	if searchErr := tx.Joins("inner join auths on users.user_id = auths.user_id").
		Where("users.gender != ? and users.area = ? and auths.reviewed_at is not null order by rand() limit ? ", user.Gender, user.Area, 3).
		Find(&users).Error; searchErr != nil {
		return errors.New("매칭 정보 찾기 실패")
	}

	if len(users) == 0 {
		return errors.New("매칭 대상이 없습니다")
	}

	var suggestions [3]models.Suggestion
	for index, user := range users {
		suggestions[index] = models.Suggestion{
			UserID:    userID,
			PartnerID: user.UserID,
		}
	}

	if err := tx.Create(&suggestions).Error; err != nil {
		return errors.New("새 매치 데이터 생성 실패")
	}
	return nil
}
