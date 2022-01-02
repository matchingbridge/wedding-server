package auth

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"wedding/datastore/mysql"
	"wedding/datastore/s3"
	"wedding/models"
)

type SignUpRequest struct {
	User     models.User           `form:"user"          json:"user"          binding:"required"`
	Password string                `form:"password"      json:"password"      binding:"required,min=8,max=16"`
	Face1    *multipart.FileHeader `form:"face1"         json:"face1"         binding:"required"`
	Face2    *multipart.FileHeader `form:"face2"         json:"face2"         binding:"required"`
	Body1    *multipart.FileHeader `form:"body1"         json:"body1"         binding:"required"`
	Body2    *multipart.FileHeader `form:"body2"         json:"body2"         binding:"required"`
	Video    *multipart.FileHeader `form:"video"         json:"video"         binding:"required"`
}

func SignUp(c *gin.Context) {
	var request SignUpRequest
	var err error

	if err = c.ShouldBind(&request); err != nil {
		err = validateSignUp(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	tx := mysql.DB.Begin()
	if err = tx.Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "데이터 베이스 연결 실패")
		return
	}
	defer tx.Rollback()

	if err = tx.Create(&request.User).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "유저 데이터 생성 실패")
		return
	}

	hashedPassword, err := hashAndSalt(request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "패스워드 확인 실패")
		return
	}

	if _, err = uploadMedia(request.User.UserID, "face1", request.Face1); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "얼굴 사진 1 업로드 실패")
		return
	}

	if _, err = uploadMedia(request.User.UserID, "face2", request.Face2); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "얼굴 사진 2 업로드 실패")
		return
	}

	if _, err = uploadMedia(request.User.UserID, "body1", request.Body1); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "전신 사진 1업로드 실패")
		return
	}

	if _, err = uploadMedia(request.User.UserID, "body2", request.Body2); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "전신 사진 2업로드 실패")
		return
	}

	if _, err = uploadMedia(request.User.UserID, "video", request.Video); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "영상 업로드 실패")
		return
	}

	newAuth := models.Auth{
		UserID:       request.User.UserID,
		Password:     hashedPassword,
		RefreshToken: newRefreshToken(),
	}
	if err = tx.Create(&newAuth).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	tx.Commit()

	c.Status(http.StatusOK)
}

func validateSignUp(err error) error { // TODO: Check with nested
	switch errType := err.(type) {
	case validator.ValidationErrors:
		for _, err := range errType {
			switch err.Field() {
			case "user":
				return errors.New("회원정보를 다시 확인해 주세요")
			case "password":
				return errors.New("비밀번호를 다시 확인해 주세요")
			case "face1":
				return errors.New("첫번째 얼굴 사진을 다시 확인해 주세요")
			case "face2":
				return errors.New("두번째 얼굴 사진을 다시 확인해 주세요")
			case "body1":
				return errors.New("첫번째 전신 사진을 다시 확인해 주세요")
			case "body2":
				return errors.New("두번째 전신 사진을 다시 확인해 주세요")
			case "video":
				return errors.New("영상을 다시 확인해 주세요")
			default:
				return errors.New(err.Field())
			}
		}
	default:
		return errors.New("field error")
	}
	return nil
}

func uploadMedia(userID string, bucket string, header *multipart.FileHeader) (key string, err error) {
	file, err := header.Open()
	if err != nil {
		return
	}
	defer file.Close()

	key = s3.MakeKey(bucket, userID)
	defer func() {
		if r := recover(); r != nil {
			s3.DeleteFile(key)
			panic(r)
		}
	}()
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = s3.UploadFile(key, fileByte)
	return
}

func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func newRefreshToken() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
