package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo_app/internal/domain/entity"
	"todo_app/pkg/utils"
)

func AuthenticationMiddleware(userRepo entity.UserRepository, userDetailRepo entity.UserDetailRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerTokenBearer := c.Request.Header.Get("Authorization")

		if headerTokenBearer == "" || !strings.HasPrefix(headerTokenBearer, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildErrorResponse(http.StatusUnauthorized, "unauthorized"))
			return
		}

		parts := strings.Split(headerTokenBearer, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildErrorResponse(http.StatusUnauthorized, "unauthorized"))
			return
		}
		token := parts[1]

		_, err := userDetailRepo.FindByToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildErrorResponse(http.StatusUnauthorized, "unauthorized"))
			return
		}

		var userPayload entity.User
		errDecrypt := utils.DecryptToken(token, &userPayload)
		if errDecrypt != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildErrorResponse(http.StatusUnauthorized, "unauthorized"))
			return
		}

		user, err := userRepo.FindByID(c, userPayload.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildErrorResponse(http.StatusUnauthorized, "unauthorized"))
			return
		}

		userContext := &entity.User{
			Username: user.Username,
			ID:       user.ID,
		}

		c.Set("userContext", userContext)
		c.Next()
	}
}
