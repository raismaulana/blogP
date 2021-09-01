package categoryapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/createcategory"
	"github.com/raismaulana/blogP/usecase/deletecategory"
	"github.com/raismaulana/blogP/usecase/showallcategories"
	"github.com/raismaulana/blogP/usecase/showcategorybyid"
	"github.com/raismaulana/blogP/usecase/updatecategory"
)

type Controller struct {
	JWTToken                *auth.JWTToken
	Env                     *envconfig.EnvConfig
	Router                  gin.IRouter
	CreateCategoryInport    createcategory.Inport
	ShowAllCategoriesInport showallcategories.Inport
	ShowCategoryByIDInport  showcategorybyid.Inport
	DeleteCategoryInport    deletecategory.Inport
	UpdateCategoryInport    updatecategory.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/v1/categories", r.authorized(), r.createCategoryHandler(r.CreateCategoryInport))
	r.Router.GET("/v1/categories", r.authorized(), r.showAllCategoriesHandler(r.ShowAllCategoriesInport))
	r.Router.GET("/v1/categories/:id_category", r.authorized(), r.showCategoryByIDHandler(r.ShowCategoryByIDInport))
	r.Router.DELETE("/v1/categories/:id_category", r.authorized(), r.deleteCategoryHandler(r.DeleteCategoryInport))
	r.Router.PUT("/v1/categories/:id_category", r.authorized(), r.updateCategoryHandler(r.UpdateCategoryInport))
}
