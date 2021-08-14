package restapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/usecase/createuser"
	"github.com/raismaulana/blogP/usecase/deleteuser"
	"github.com/raismaulana/blogP/usecase/showallusers"
	"github.com/raismaulana/blogP/usecase/showuserbyemail"
	"github.com/raismaulana/blogP/usecase/showuserbyid"
	"github.com/raismaulana/blogP/usecase/showuserbyusername"
	"github.com/raismaulana/blogP/usecase/updateuser"
)

type Controller struct {
	Router                   gin.IRouter
	CreateUserInport         createuser.Inport
	ShowUserByIDInport       showuserbyid.Inport
	ShowUserByEmailInport    showuserbyemail.Inport
	ShowUserByUsernameInport showuserbyusername.Inport
	ShowAllUsersInport       showallusers.Inport
	UpdateUserInport         updateuser.Inport
	DeleteUserInport         deleteuser.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/user/create", r.authorized(), r.CreateUserHandler(r.CreateUserInport))
	r.Router.GET("/user/id/:id_user", r.authorized(), r.showUserByIDHandler(r.ShowUserByIDInport))
	r.Router.GET("/user/email/:email", r.authorized(), r.showUserByEmailHandler(r.ShowUserByEmailInport))
	r.Router.GET("/user/username/:username", r.authorized(), r.showUserByUsernameHandler(r.ShowUserByUsernameInport))
	r.Router.GET("/user/all", r.authorized(), r.showAllUsersHandler(r.ShowAllUsersInport))
	r.Router.PUT("/user/update", r.authorized(), r.updateUserHandler(r.UpdateUserInport))
	r.Router.DELETE("/user/delete/:id_user", r.authorized(), r.deleteUserHandler(r.DeleteUserInport))
}
