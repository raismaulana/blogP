package registry

import (
	"context"
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/blogP/application"
	"github.com/raismaulana/blogP/controller/categoryapi"
	"github.com/raismaulana/blogP/controller/postapi"
	"github.com/raismaulana/blogP/controller/tagapi"
	"github.com/raismaulana/blogP/controller/userapi"
	"github.com/raismaulana/blogP/gateway/master"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/server"
	"github.com/raismaulana/blogP/usecase/activationuser"
	"github.com/raismaulana/blogP/usecase/createcategory"
	"github.com/raismaulana/blogP/usecase/createpost"
	"github.com/raismaulana/blogP/usecase/createtag"
	"github.com/raismaulana/blogP/usecase/createuser"
	"github.com/raismaulana/blogP/usecase/deletecategory"
	"github.com/raismaulana/blogP/usecase/deletepost"
	"github.com/raismaulana/blogP/usecase/deletetag"
	"github.com/raismaulana/blogP/usecase/deleteuser"
	"github.com/raismaulana/blogP/usecase/forgotpassword"
	"github.com/raismaulana/blogP/usecase/loginuser"
	"github.com/raismaulana/blogP/usecase/resetactivationuser"
	"github.com/raismaulana/blogP/usecase/showallcategories"
	"github.com/raismaulana/blogP/usecase/showallposts"
	"github.com/raismaulana/blogP/usecase/showalltags"
	"github.com/raismaulana/blogP/usecase/showalluserposts"
	"github.com/raismaulana/blogP/usecase/showallusers"
	"github.com/raismaulana/blogP/usecase/showcategorybyid"
	"github.com/raismaulana/blogP/usecase/showpostbyid"
	"github.com/raismaulana/blogP/usecase/showpostbyslug"
	"github.com/raismaulana/blogP/usecase/showtagbyid"
	"github.com/raismaulana/blogP/usecase/showuserbyemail"
	"github.com/raismaulana/blogP/usecase/showuserbyid"
	"github.com/raismaulana/blogP/usecase/showuserbyusername"
	"github.com/raismaulana/blogP/usecase/updatecategory"
	"github.com/raismaulana/blogP/usecase/updatepassword"
	"github.com/raismaulana/blogP/usecase/updatepost"
	"github.com/raismaulana/blogP/usecase/updatetag"
	"github.com/raismaulana/blogP/usecase/updateuser"
	"github.com/raismaulana/blogP/usecase/uploaduserphotoprofile"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type usingdb struct {
	server.GinHTTPHandler
	categoryapiController categoryapi.Controller
	postapiController     postapi.Controller
	userapiController     userapi.Controller
	tagapiController      tagapi.Controller
	// TODO Another controller will added here ... <<<<<<
}

var (
	ctx = context.Background()
)

func NewUsingdb() func() application.RegistryContract {
	return func() application.RegistryContract {

		env := setupEnv()
		jwtToken := setupJWTToken(env)
		db := setupDB(env)
		rdb := setupRedis(env)
		httpHandler := setupHTTPHandler(env)
		enforcer := setupCasbinEnforcer()

		datasource, err := master.NewMasterGateway(env, db, rdb, jwtToken)
		if err != nil {
			log.Error(ctx, "%v", err.Error())
			os.Exit(1)
		}

		return &usingdb{
			GinHTTPHandler: httpHandler,
			categoryapiController: categoryapi.Controller{
				JWTToken:                jwtToken,
				Env:                     env,
				Enforcer:                enforcer,
				Router:                  httpHandler.Router,
				CreateCategoryInport:    createcategory.NewUsecase(datasource),
				ShowAllCategoriesInport: showallcategories.NewUsecase(datasource),
				ShowCategoryByIDInport:  showcategorybyid.NewUsecase(datasource),
				DeleteCategoryInport:    deletecategory.NewUsecase(datasource),
				UpdateCategoryInport:    updatecategory.NewUsecase(datasource),
			},
			postapiController: postapi.Controller{
				JWTToken:             jwtToken,
				Env:                  env,
				Enforcer:             enforcer,
				Router:               httpHandler.Router,
				CreatePostInport:     createpost.NewUsecase(datasource),
				ShowAllPostsInport:   showallposts.NewUsecase(datasource),
				ShowPostBySlugInport: showpostbyslug.NewUsecase(datasource),
				ShowPostByIDInport:   showpostbyid.NewUsecase(datasource),
				DeletePostInport:     deletepost.NewUsecase(datasource),
				UpdatePostInport:     updatepost.NewUsecase(datasource),
			},
			userapiController: userapi.Controller{
				JWTToken:                     jwtToken,
				Env:                          env,
				Enforcer:                     enforcer,
				Router:                       httpHandler.Router,
				CreateUserInport:             createuser.NewUsecase(datasource),
				ShowUserByIDInport:           showuserbyid.NewUsecase(datasource),
				ShowUserByEmailInport:        showuserbyemail.NewUsecase(datasource),
				ShowUserByUsernameInport:     showuserbyusername.NewUsecase(datasource),
				ShowAllUsersInport:           showallusers.NewUsecase(datasource),
				UpdateUserInport:             updateuser.NewUsecase(datasource),
				DeleteUserInport:             deleteuser.NewUsecase(datasource),
				ActivationUserInport:         activationuser.NewUsecase(datasource),
				ResetActivationUserInport:    resetactivationuser.NewUsecase(datasource),
				LoginUserInport:              loginuser.NewUsecase(datasource),
				ShowAllUserPostsInport:       showalluserposts.NewUsecase(datasource),
				UpdatePasswordInport:         updatepassword.NewUsecase(datasource),
				ForgotPasswordInport:         forgotpassword.NewUsecase(datasource),
				UploadUserPhotoProfileInport: uploaduserphotoprofile.NewUsecase(datasource),
			},
			tagapiController: tagapi.Controller{
				JWTToken:          jwtToken,
				Env:               env,
				Enforcer:          enforcer,
				Router:            httpHandler.Router,
				ShowAllTagsInport: showalltags.NewUsecase(datasource),
				CreateTagInport:   createtag.NewUsecase(datasource),
				ShowTagByIDInport: showtagbyid.NewUsecase(datasource),
				UpdateTagInport:   updatetag.NewUsecase(datasource),
				DeleteTagInport:   deletetag.NewUsecase(datasource),
			},
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *usingdb) SetupController() {
	r.userapiController.RegisterRouter()
	r.postapiController.RegisterRouter()
	r.tagapiController.RegisterRouter()
	r.categoryapiController.RegisterRouter()

	// TODO another router call will added here ... <<<<<<
}

func setupEnv() *envconfig.EnvConfig {
	env, err := envconfig.NewEnvConfig()
	if err != nil {
		log.Error(ctx, "Config Problem %v", err.Error())
		os.Exit(1)
	}
	return env
}

func setupJWTToken(env *envconfig.EnvConfig) *auth.JWTToken {
	jwtToken, err := auth.NewJWTToken(env)
	if err != nil {
		log.Error(context.Background(), "Secret Key Problem %v", err.Error())
		os.Exit(1)
	}
	return jwtToken
}

func setupDB(env *envconfig.EnvConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func setupRedis(env *envconfig.EnvConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost + env.RedisPort,
		Password: env.RedisPassword,
		DB:       env.RedisDB,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error(ctx, "%v", err.Error())
		os.Exit(1)
	}
	return rdb
}

func setupHTTPHandler(env *envconfig.EnvConfig) server.GinHTTPHandler {
	httpHandler, err := server.NewGinHTTPHandler(":" + env.AppPort)
	if err != nil {
		log.Error(ctx, "%v", err.Error())
		os.Exit(1)
	}
	return httpHandler
}

func setupCasbinEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("infrastructure/auth/casbin_model.conf", "infrastructure/auth/casbin_policy.csv")
	if err != nil {
		log.Error(ctx, "%v", err.Error())
		os.Exit(1)
	}
	return e
}
