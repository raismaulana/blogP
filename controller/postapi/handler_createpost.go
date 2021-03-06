package postapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"github.com/raismaulana/blogP/usecase/createpost"
)

// createPostHandler ...
func (r *Controller) createPostHandler(inputPort createpost.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var req createpost.InportRequest
		req.UserID = c.MustGet("auth_id_user").(int64)
		if err := c.BindJSON(&req); err != nil {
			log.Error(ctx, "bind", err.Error())
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.FailUnmarshalResponseBodyError))
				return
			}
			var ez string
			for _, e := range errs {
				ez = ez + e.Translate(util.Trans) + "\n"
			}

			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.ValidationError.Var(ez)))
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
