package postapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
)

// authorized is an interceptor
func (r *Controller) authorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		authorization := c.GetHeader("Authorization")
		encodedToken := strings.TrimPrefix(authorization, "Bearer ")
		if authorization == encodedToken {
			log.Error(ctx, "failed getting bearer", encodedToken)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		token, err := r.JWTToken.VerifyToken(encodedToken)
		if err != nil {
			log.Error(ctx, "failed verify token", token)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		claims := token.Claims.(*auth.CustomClaims)

		if claims, ok := token.Claims.(*auth.CustomClaims); !(ok && token.Valid) {
			log.Error(ctx, "token is not valid", claims)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		log.Info(ctx, util.MustJSON(claims))

		authorized := claims.Activated
		c.Set("auth_id_user", claims.ID)
		c.Set("auth_subject", claims.Subject)
		c.Set("auth_role", claims.Role)
		c.Set("auth_email", claims.Email)
		c.Set("auth_activated", claims.Activated)

		if !authorized {
			log.Error(ctx, "unauthorized", claims)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
