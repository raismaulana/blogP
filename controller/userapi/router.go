package userapi

import (
	"github.com/gin-gonic/gin"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/usecase/activationuser"
	"github.com/raismaulana/blogP/usecase/createuser"
	"github.com/raismaulana/blogP/usecase/deleteuser"
	"github.com/raismaulana/blogP/usecase/forgotpassword"
	"github.com/raismaulana/blogP/usecase/loginuser"
	"github.com/raismaulana/blogP/usecase/resetactivationuser"
	"github.com/raismaulana/blogP/usecase/showalluserposts"
	"github.com/raismaulana/blogP/usecase/showallusers"
	"github.com/raismaulana/blogP/usecase/showuserbyemail"
	"github.com/raismaulana/blogP/usecase/showuserbyid"
	"github.com/raismaulana/blogP/usecase/showuserbyusername"
	"github.com/raismaulana/blogP/usecase/updatepassword"
	"github.com/raismaulana/blogP/usecase/updateuser"
)

type Controller struct {
	JWTToken                  *auth.JWTToken
	Env                       *envconfig.EnvConfig
	Router                    gin.IRouter
	CreateUserInport          createuser.Inport
	ShowUserByIDInport        showuserbyid.Inport
	ShowUserByEmailInport     showuserbyemail.Inport
	ShowUserByUsernameInport  showuserbyusername.Inport
	ShowAllUsersInport        showallusers.Inport
	UpdateUserInport          updateuser.Inport
	DeleteUserInport          deleteuser.Inport
	ActivationUserInport      activationuser.Inport
	ResetActivationUserInport resetactivationuser.Inport
	LoginUserInport           loginuser.Inport
	ShowAllUserPostsInport    showalluserposts.Inport
	UpdatePasswordInport      updatepassword.Inport
	ForgotPasswordInport      forgotpassword.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.DELETE("/v1/users/:id_user", r.authorized(), r.deleteUserHandler(r.DeleteUserInport))
	r.Router.GET("/v1/users", r.showAllUsersHandler(r.ShowAllUsersInport))
	r.Router.GET("/v1/users/:id_user", r.showUserByIDHandler(r.ShowUserByIDInport))
	r.Router.GET("/v1/users/:id_user/activation", r.activationUserHandler(r.ActivationUserInport))
	r.Router.GET("/v1/users/email/:email", r.showUserByEmailHandler(r.ShowUserByEmailInport))
	r.Router.GET("/v1/users/username/:username", r.showUserByUsernameHandler(r.ShowUserByUsernameInport))
	r.Router.GET("/v1/users/username/:username/posts", r.showAllUserPostsHandler(r.ShowAllUserPostsInport))
	r.Router.POST("/v1/users", r.CreateUserHandler(r.CreateUserInport))
	r.Router.PUT("/v1/users/:id_user", r.authorized(), r.updateUserHandler(r.UpdateUserInport))
	r.Router.GET("/v1/users/:id_user/re-activation", r.authorized(), r.resetActivationUserHandler(r.ResetActivationUserInport)) // broken
	r.Router.POST("/v1/users/auth", r.loginUserHandler(r.LoginUserInport))
	r.Router.PATCH("/v1/users/:id_user/password", r.authorized(), r.updatePasswordHandler(r.UpdatePasswordInport))
	r.Router.POST("/v1/users/forgotpassword", r.forgotPasswordHandler(r.ForgotPasswordInport))
}
