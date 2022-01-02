package auth

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"wedding/datastore/mysql"
	"wedding/models"
)

type SignInRequest struct {
	Email    string `form:"email"  json:"email"  binding:"required,min=8,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

func SignIn(c *gin.Context) {
	var request SignInRequest
	if err := validateSignIn(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := mysql.DB.Where("email = ?", request.Email).Take(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "유저 정보를 찾을 수 없습니다")
		return
	}

	var auth models.Auth
	if err := mysql.DB.Where("user_id = ?", user.UserID).Take(&auth).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, "가입 정보를 찾을 수 없습니다")
		return
	}

	if !compare(auth.Password, request.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, "비밀번호가 일치하지 않습니다")
		return
	}

	if auth.ReviewedAt == nil {
		c.JSON(http.StatusAccepted, "프로필 리뷰중입니다.")
		return
	}

	accessToken, expiresAt, err := createAccessToken(auth.AuthID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "다시 로그인해주세요")
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "expires_at": expiresAt, "refresh_token": auth.RefreshToken})
}

func validateSignIn(c *gin.Context, request *SignInRequest) error {
	if validationErr := c.ShouldBindJSON(&request); validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			if err.Field() == "user_id" {
				return errors.New("아이디를 다시 확인해 주세요")
			}

			if err.Field() == "password" {
				return errors.New("비밀번호를 다시 확인해 주세요")
			}
		}
	}
	return nil
}

func compare(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}

func createAccessToken(authID int) (ss string, expiresAt int64, err error) {
	expiresAt = time.Now().Add(24 * time.Hour).Unix()
	claims := &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Subject:   strconv.Itoa(authID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString([]byte("wedding"))
	return
}
