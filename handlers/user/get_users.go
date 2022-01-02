package user

import (
	"net/http"
	"time"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	UserID     string     `json:"user_id"`
	AuthID     int        `json:"auth_id"`
	Name       string     `json:"name"`
	Gender     int        `json:"gender"`
	Marriage   int        `json:"marriage"`
	Area       int        `json:"area"`
	Scholar    int        `json:"scholar"`
	Job        int        `json:"job"`
	ReviewedAt *time.Time `json:"reviewed_at"`
}

func GetUsers(c *gin.Context) {
	var results []UserAuth
	mysql.DB.Model(&models.User{}).Select("users.user_id, auths.auth_id, users.name, users.gender, users.marriage, users.area, users.scholar, users.job, auths.reviewed_at").Joins("left join auths on auths.user_id = users.user_id").Scan(&results)

	c.JSON(http.StatusOK, results)
}
