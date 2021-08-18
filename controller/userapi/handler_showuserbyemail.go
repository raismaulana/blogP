package userapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"github.com/raismaulana/blogP/usecase/showuserbyemail"
)

// showUserByEmailHandler ...
func (r *Controller) showUserByEmailHandler(inputPort showuserbyemail.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		req := showuserbyemail.InportRequest{
			Email: c.Param("email"),
		}

		log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusOK, NewErrorResponse(err))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
