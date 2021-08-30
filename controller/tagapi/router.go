package tagapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/showalltags"
)

type Controller struct {
	JWTToken          *auth.JWTToken
	Env               *envconfig.EnvConfig
	Router            gin.IRouter
	ShowAllTagsInport showalltags.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("v1/tags", r.authorized(), r.showAllTagsHandler(r.ShowAllTagsInport))
}
