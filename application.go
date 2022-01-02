package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"wedding/datastore/mysql"
	"wedding/datastore/s3"
	authHandler "wedding/handlers/auth"
	baseHandler "wedding/handlers/base"
	chatHandler "wedding/handlers/chat"
	detailHandler "wedding/handlers/detail"
	matchHandler "wedding/handlers/match"
	suggestionHandler "wedding/handlers/suggestion"
	userHandler "wedding/handlers/user"
)

func main() {
	mysql.ConnectDB()
	mysql.CreateTable()

	s3.ConnectBucket()

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	setupHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5206"
	}

	server := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}

	log.Println("Starting Server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func setupHandler(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	client := r.Group("/api/client")
	{
		auth := client.Group("/auth")
		{
			auth.POST("/signin", authHandler.SignIn)
			auth.POST("/signup", authHandler.SignUp)
		}
		chat := client.Group("/chat")
		{
			chat.GET("", baseHandler.UserAuthentication, chatHandler.GetChat)
			chat.POST("", baseHandler.UserAuthentication, chatHandler.MakeChat)
		}
		detail := client.Group("/detail")
		{
			detail.GET("", baseHandler.UserAuthentication, detailHandler.GetDetail)
			detail.POST("", baseHandler.UserAuthentication, detailHandler.UpdateDetail)
		}
		match := client.Group("/match")
		{
			match.GET("", baseHandler.UserAuthentication, matchHandler.GetMatch)
			match.PUT(fmt.Sprintf("/ask/:%s", "partner_id"), baseHandler.UserAuthentication, matchHandler.AskMatch)
			match.PUT(fmt.Sprintf("/answer/:%s", "partner_id"), baseHandler.UserAuthentication, matchHandler.AnswerMatch)
			match.PUT(fmt.Sprintf("/terminate/:%s", "partner_id"), baseHandler.UserAuthentication, matchHandler.TerminateMatch)
			match.GET("/matched", baseHandler.UserAuthentication, matchHandler.GetMatched)
			match.GET("/unmatched", baseHandler.UserAuthentication, matchHandler.GetUnmatched)
		}
		suggestion := client.Group("/suggestion")
		{
			suggestion.GET("", baseHandler.UserAuthentication, suggestionHandler.GetSuggestions)
		}
		user := client.Group("/user")
		{
			user.GET("", baseHandler.UserAuthentication, userHandler.GetUser)
			user.GET(fmt.Sprintf("/email/:%s", "email"), userHandler.CheckEmail)
			user.GET(fmt.Sprintf("/media/:%s/:%s", "bucket", "user_id"), userHandler.GetMedia)
			user.GET(fmt.Sprintf(":%s", "partner_id"), baseHandler.UserAuthentication, userHandler.GetPartner)
		}
	}

	admin := r.Group("/api/admin")
	{
		auth := admin.Group("/auth")
		{
			auth.OPTIONS(fmt.Sprintf("/review/:%s", "auth_id"), baseHandler.AdminAuthentication, func(c *gin.Context) {
				c.AbortWithStatus(http.StatusNoContent)
			})
			auth.PUT(fmt.Sprintf("/review/:%s", "auth_id"), baseHandler.AdminAuthentication, authHandler.Review)
		}
		chat := admin.Group("/chat")
		{
			chat.GET("", baseHandler.AdminAuthentication, chatHandler.GetChats)
			chat.PUT(fmt.Sprintf("/:%s", "user_id"), baseHandler.AdminAuthentication, chatHandler.ReplyChat)
		}
		match := admin.Group("/match")
		{
			match.GET("", baseHandler.AdminAuthentication, matchHandler.GetMatches)
		}
		user := admin.Group("/user")
		{
			user.GET("", baseHandler.AdminAuthentication, userHandler.GetUsers)
		}
	}
}
