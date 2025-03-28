// auth/auth.go

package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/utils"
)

var salt = config.Salt
var secretKey = []byte(salt)

// VerifyPassword verifies if the provided password matches the hashed password

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"id"`
}

func VerifyToken(bearer_token string) (string, string, error) {
	token, err := jwt.ParseWithClaims(bearer_token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", "invalid-token", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", "invalid-token", nil
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return "", "expired", nil
	}

	return claims.UserID, "valid", nil
}

func AuthUser(c *gin.Context) {
	var currAccData models.AccountData
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			token := parts[1]
			currAccData.UserID, currAccData.VerifyStatus, currAccData.ErrVerif = VerifyToken(token)

			if currAccData.VerifyStatus == "invalid-token" || currAccData.VerifyStatus == "expired" {
				currAccData.UserID = ""
				utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "Your session is expired, Please re-Login!"})
				c.Abort()
				return
			}

			c.Set("accountData", currAccData)
			c.Next()
		} else {
			currAccData.UserID = ""
			currAccData.VerifyStatus = "invalid-token-format"
			currAccData.ErrVerif = nil
			utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "Invalid token format. Please check your Authorization header."})
			c.Abort()
			return
		}
	} else {
		currAccData.UserID = ""
		currAccData.VerifyStatus = "no-token"
		currAccData.ErrVerif = nil
		utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "You have to login first!"})
		c.Abort()
		return
	}
}
