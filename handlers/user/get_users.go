package user

import (
	"net/http"
	"time"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

type UserAuth struct {
	UserID string `json:"user_id"`
	AuthID int    `json:"auth_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`

	Area      *int  `json:"area"`
	Gender    *int  `json:"gender"`
	Marriage  *int  `json:"marriage"`
	Height    *int  `json:"height"`
	Weight    *int  `json:"weight"`
	Smoke     *bool `json:"smoke"`
	Drink     *bool `json:"drink"`
	Religion  *int  `json:"religion"`
	BodyType  *int  `json:"bodytype"`
	Character *int  `json:"character"`

	Scholar *int `json:"scholar"`
	Job     *int `json:"job"`

	ReviewedAt *time.Time `json:"reviewed_at"`
}

func GetUsers(c *gin.Context) {
	var results []UserAuth
	mysql.DB.Model(&models.User{}).Select("users.user_id, auths.auth_id, users.email, users.name, users.phone, users.area, users.gender, users.marriage, users.height, users.weight, users.smoke, users.drink, users.religion, users.body_type, users.character, users.scholar, users.job, auths.reviewed_at").Joins("left join auths on auths.user_id = users.user_id").Scan(&results)

	c.JSON(http.StatusOK, results)
}
