package user

import (
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func CheckEmail(c *gin.Context) {
	email := c.Param("email")

	var user models.User
	result := mysql.DB.Where("email = ?", email).Take(&user)
	c.JSON(http.StatusOK, result.RowsAffected == 0)
}
