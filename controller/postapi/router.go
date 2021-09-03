package postapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/createpost"
)

type Controller struct {
	JWTToken         *auth.JWTToken
	Env              *envconfig.EnvConfig
	Router           gin.IRouter
	CreatePostInport createpost.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/v1/posts", r.authorized(), r.createPostHandler(r.CreatePostInport))
}
