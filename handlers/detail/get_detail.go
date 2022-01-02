package detail

import (
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {
	userID := c.GetString("user_id")

	var detail models.Detail
	query := "user_id = ?"
	if err := mysql.DB.Where(query, userID).Take(&detail).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, detail)
}
