package main

import (
	"backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthUser struct {
	UserId   string
	Email    string
	UserName string
}

// a simple middleware to verify JWT Access Token
func AuthMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.Request.Header.Get("Authorization")

		token, ok := strings.CutPrefix(authHeader, "Bearer ")

		if !ok {
			g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization is missing"})
			return
		}

		user, err := utils.ParseAccessToken(token)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token received"})
			return
		}

		db, err := GetDB(DbPath)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to establish connection to database"})
			return
		}

		query, err := db.Prepare(`select user_name from Users where ROWID = ? and email = ?`)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate database statement"})
			return
		}

		defer query.Close()

		var user_name string

		err = query.QueryRow(user.UserId, user.Email).Scan(&user_name)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "UserID was not found"})
			return
		}

		g.Set("User", AuthUser{UserId: user.UserId, Email: user.Email, UserName: user_name})

		g.Next()
	}
}
