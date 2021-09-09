package categoryapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
)

// authorized is an interceptor
func (r *Controller) authorized() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.Set("auth_id_user", 0)
			c.Set("auth_subject", c.ClientIP())
			c.Set("auth_role", "guest")
			c.Set("auth_email", "")
			c.Set("auth_activated", false)
		} else {

			encodedToken := strings.TrimPrefix(authorization, "Bearer ")
			if authorization == encodedToken {
				log.Error(ctx, "failed getting bearer", encodedToken)
			}

			token, err := r.JWTToken.VerifyToken(encodedToken)
			if err != nil {
				log.Error(ctx, err.Error(), token)
				c.AbortWithStatusJSON(401, NewErrorResponse(apperror.AuthenticationError.Var(err.Error())))
				return
			} else {

				claims := token.Claims.(*auth.CustomClaims)
				log.Info(ctx, util.MustJSON(claims))

				if claims, ok := token.Claims.(*auth.CustomClaims); !(ok && token.Valid) {
					log.Error(ctx, "token is not valid", claims)
					c.AbortWithStatusJSON(401, NewErrorResponse(apperror.AuthenticationError.Var("token is not valid")))
					return
				} else {
					c.Set("auth_id_user", claims.ID)
					c.Set("auth_subject", claims.Subject)
					c.Set("auth_role", claims.Role)
					c.Set("auth_email", claims.Email)
					c.Set("auth_activated", claims.Activated)

				}
			}
		}

		ok, err := r.Enforcer.Enforce(c.MustGet("auth_role").(string), c.Request.URL.Path, c.Request.Method)

		if err != nil {
			log.Error(ctx, err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !ok {
			log.Error(ctx, "unauthorized", c.MustGet("auth_role").(string), c.Request.URL.Path, c.Request.Method)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		log.Info(ctx, "authorized", c.MustGet("auth_role").(string), c.Request.URL.Path, c.Request.Method)

	}
}
