package tagapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/createtag"
	"github.com/raismaulana/blogP/usecase/deletetag"
	"github.com/raismaulana/blogP/usecase/showalltags"
	"github.com/raismaulana/blogP/usecase/showtagbyid"
	"github.com/raismaulana/blogP/usecase/updatetag"
)

type Controller struct {
	JWTToken          *auth.JWTToken
	Env               *envconfig.EnvConfig
	Router            gin.IRouter
	ShowAllTagsInport showalltags.Inport
	CreateTagInport   createtag.Inport
	ShowTagByIDInport showtagbyid.Inport
	UpdateTagInport   updatetag.Inport
	DeleteTagInport   deletetag.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/v1/tags", r.showAllTagsHandler(r.ShowAllTagsInport))
	r.Router.POST("/v1/tags", r.authorized(), r.createTagHandler(r.CreateTagInport))
	r.Router.GET("v1/tags/:id_tag", r.showTagByIDHandler(r.ShowTagByIDInport))
	r.Router.PUT("/v1/tags/:id_tag", r.authorized(), r.updateTagHandler(r.UpdateTagInport))
	r.Router.DELETE("/v1/tags/:id_tag", r.authorized(), r.deleteTagHandler(r.DeleteTagInport))
}
