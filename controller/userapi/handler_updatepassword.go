package userapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"github.com/raismaulana/blogP/usecase/updatepassword"
)

// updatePasswordHandler ...
func (r *Controller) updatePasswordHandler(inputPort updatepassword.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		id, err := strconv.ParseInt(c.Param("id_user"), 10, 64)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.NumberOnlyParam))
			return
		}

		var req updatepassword.InportRequest
		req.ID = id
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
