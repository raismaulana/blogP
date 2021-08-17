package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/activationuser"
	"github.com/raismaulana/blogP/usecase/createuser"
	"github.com/raismaulana/blogP/usecase/deleteuser"
	"github.com/raismaulana/blogP/usecase/showallusers"
	"github.com/raismaulana/blogP/usecase/showuserbyemail"
	"github.com/raismaulana/blogP/usecase/showuserbyid"
	"github.com/raismaulana/blogP/usecase/showuserbyusername"
	"github.com/raismaulana/blogP/usecase/updateuser"
)

type Controller struct {
	Env                      *envconfig.EnvConfig
	Router                   gin.IRouter
	CreateUserInport         createuser.Inport
	ShowUserByIDInport       showuserbyid.Inport
	ShowUserByEmailInport    showuserbyemail.Inport
	ShowUserByUsernameInport showuserbyusername.Inport
	ShowAllUsersInport       showallusers.Inport
	UpdateUserInport         updateuser.Inport
	DeleteUserInport         deleteuser.Inport
	ActivationUserInport     activationuser.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.DELETE("/users/:id_user", r.authorized(), r.deleteUserHandler(r.DeleteUserInport))
	r.Router.GET("/users", r.authorized(), r.showAllUsersHandler(r.ShowAllUsersInport))
	r.Router.GET("/users/:id_user", r.authorized(), r.showUserByIDHandler(r.ShowUserByIDInport))
	r.Router.GET("/users/:id_user/activation", r.authorized(), r.activationUserHandler(r.ActivationUserInport))
	r.Router.GET("/users/email/:email", r.authorized(), r.showUserByEmailHandler(r.ShowUserByEmailInport))
	r.Router.GET("/users/username/:username", r.authorized(), r.showUserByUsernameHandler(r.ShowUserByUsernameInport))
	r.Router.POST("/users", r.authorized(), r.CreateUserHandler(r.CreateUserInport))
	r.Router.PUT("/users/:id_user", r.authorized(), r.updateUserHandler(r.UpdateUserInport))
}
