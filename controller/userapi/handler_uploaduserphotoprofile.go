package userapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raismaulana/blogP/application/apperror"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/util"
	"github.com/raismaulana/blogP/usecase/uploaduserphotoprofile"
)

// uploadUserPhotoProfileHandler ...
func (r *Controller) uploadUserPhotoProfileHandler(inputPort uploaduserphotoprofile.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		id, err := strconv.ParseInt(c.Param("id_user"), 10, 64)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.NumberOnlyParam))
			return
		}

		var req uploaduserphotoprofile.InportRequest
		req.ID = id
		if err := c.ShouldBind(&req); err != nil {
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
		if req.PhotoProfile.Size > 1048576 {
			log.Error(ctx, "size: "+strconv.FormatInt(req.PhotoProfile.Size, 10))
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.FileAllowedMaxSizeIs.Var("1MB")))
			return
		}
		contenttype := req.PhotoProfile.Header.Get("Content-Type")
		if contenttype != "image/jpeg" {
			log.Error(ctx, contenttype)
			c.JSON(http.StatusBadRequest, NewErrorResponse(apperror.OnlyJPEGOrJPGAllowed))
			return
		}
		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
