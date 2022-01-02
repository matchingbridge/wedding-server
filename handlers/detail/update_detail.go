package detail

import (
	"errors"
	"net/http"
	"wedding/datastore/mysql"
	"wedding/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UpdateDetailRequest struct {
	Detail models.Detail `form:"detail" json:"detail" binding:"required"`
}

func UpdateDetail(c *gin.Context) {
	userID := c.GetString("user_id")

	var request UpdateDetailRequest
	var err error

	if err = c.ShouldBind(&request); err != nil {
		err = validateUpdateDetail(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	request.Detail.UserID = userID

	if err := mysql.DB.Create(&request.Detail).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("만남 요청 실패"))
		return
	}

	c.JSON(http.StatusOK, request.Detail)
}

func validateUpdateDetail(err error) error { // TODO: Check with nested
	switch errType := err.(type) {
	case validator.ValidationErrors:
		for _, err := range errType {
			switch err.Field() {
			case "detail":
				return errors.New("상세정보를 다시 확인해 주세요")
			default:
				return errors.New(err.Field())
			}
		}
	default:
		return errors.New("field error")
	}
	return nil
}
