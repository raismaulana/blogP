package userapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"github.com/raismaulana/blogP/usecase/createuser"
)

// createuserHandler ...
func (r *Controller) CreateUserHandler(inputPort createuser.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req createuser.InportRequest
		if err := c.BindJSON(&req); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(newErr))
			return
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
