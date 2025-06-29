// auth/auth.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
	"godp.abdanhafidz.com/utils"
)

func AuthUser(c *gin.Context) {
	var currAccData models.AccountData
	if c.Request.Header["Authorization"] != nil {
		token := c.Request.Header["Authorization"]

		currAccData.UserID, currAccData.VerifyStatus, currAccData.ErrVerif = services.VerifyToken(token[0])

		if currAccData.VerifyStatus == "invalid-token" || currAccData.VerifyStatus == "expired" {
			currAccData.UserID = uuid.UUID{}
			utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "Your session is expired, Please re-Login!"})
			c.Abort()
			return
		} else {
			c.Set("accountData", currAccData)
			c.Next()
		}
	} else {
		currAccData.UserID = uuid.UUID{}
		currAccData.VerifyStatus = "no-token"
		currAccData.ErrVerif = nil
		utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "You have to login first!"})
		c.Abort()
		return
	}

}
