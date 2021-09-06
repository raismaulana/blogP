package postapi

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/createpost"
	"github.com/raismaulana/blogP/usecase/deletepost"
	"github.com/raismaulana/blogP/usecase/showallposts"
	"github.com/raismaulana/blogP/usecase/showpostbyid"
	"github.com/raismaulana/blogP/usecase/showpostbyslug"
	"github.com/raismaulana/blogP/usecase/updatepost"
)

type Controller struct {
	JWTToken             *auth.JWTToken
	Env                  *envconfig.EnvConfig
	Enforcer             *casbin.Enforcer
	Router               gin.IRouter
	CreatePostInport     createpost.Inport
	ShowAllPostsInport   showallposts.Inport
	ShowPostBySlugInport showpostbyslug.Inport
	ShowPostByIDInport   showpostbyid.Inport
	DeletePostInport     deletepost.Inport
	UpdatePostInport     updatepost.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/v1/posts", r.authorized(), r.createPostHandler(r.CreatePostInport))
	r.Router.GET("/v1/posts", r.authorized(), r.showAllPostsHandler(r.ShowAllPostsInport))
	r.Router.GET("/v1/posts/:id_post", r.authorized(), r.showPostByIDHandler(r.ShowPostByIDInport))
	r.Router.GET("/v1/posts/slug/:slug", r.authorized(), r.showPostBySlugHandler(r.ShowPostBySlugInport))
	r.Router.PUT("/v1/posts/:id_post", r.authorized(), r.updatePostHandler(r.UpdatePostInport))
	r.Router.DELETE("/v1/posts/:id_post", r.authorized(), r.deletePostHandler(r.DeletePostInport))
}
