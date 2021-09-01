package tagapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"github.com/raismaulana/blogP/usecase/deletetag"
)

// deleteTagHandler ...
func (r *Controller) deleteTagHandler(inputPort deletetag.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		id, err := strconv.ParseInt(c.Param("id_tag"), 10, 64)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.NumberOnlyParam))
			return
		}

		var req deletetag.InportRequest
		req.ID = id

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
