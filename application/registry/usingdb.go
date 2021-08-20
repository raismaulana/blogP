package registry

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/blogP/application"
	"github.com/raismaulana/blogP/controller/userapi"
	"github.com/raismaulana/blogP/gateway/master"
	"github.com/raismaulana/blogP/infrastructure/auth"
	"github.com/raismaulana/blogP/infrastructure/envconfig"
	"github.com/raismaulana/blogP/infrastructure/log"
	"github.com/raismaulana/blogP/infrastructure/server"
	"github.com/raismaulana/blogP/usecase/activationuser"
	"github.com/raismaulana/blogP/usecase/createuser"
	"github.com/raismaulana/blogP/usecase/deleteuser"
	"github.com/raismaulana/blogP/usecase/loginuser"
	"github.com/raismaulana/blogP/usecase/resetactivationuser"
	"github.com/raismaulana/blogP/usecase/showallusers"
	"github.com/raismaulana/blogP/usecase/showuserbyemail"
	"github.com/raismaulana/blogP/usecase/showuserbyid"
	"github.com/raismaulana/blogP/usecase/showuserbyusername"
	"github.com/raismaulana/blogP/usecase/updateuser"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type usingdb struct {
	server.GinHTTPHandler
	userapiController userapi.Controller
	// TODO Another controller will added here ... <<<<<<
}

func NewUsingdb() func() application.RegistryContract {
	return func() application.RegistryContract {
		ctx := context.Background()

		env, err := envconfig.NewEnvConfig()
		if err != nil {
			log.Error(ctx, "Config Problem %v", err.Error())
			os.Exit(1)
		}

		jwtToken, err := auth.NewJWTToken(env)
		if err != nil {
			log.Error(context.Background(), "Secret Key Problem %v", err.Error())
			os.Exit(1)
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			env.DBHost,
			env.DBUser,
			env.DBPassword,
			env.DBName,
			env.DBPort,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		rdb := redis.NewClient(&redis.Options{
			Addr:     env.RedisHost + ":" + env.RedisPort,
			Password: env.RedisPassword,
			DB:       env.RedisDB,
		})
		_, err = rdb.Ping(ctx).Result()
		if err != nil {
			log.Error(ctx, "%v", err.Error())
			os.Exit(1)
		}

		httpHandler, err := server.NewGinHTTPHandler(":" + env.AppPort)
		if err != nil {
			log.Error(ctx, "%v", err.Error())
			os.Exit(1)
		}

		datasource, err := master.NewMasterGateway(env, db, rdb, jwtToken)
		if err != nil {
			log.Error(ctx, "%v", err.Error())
			os.Exit(1)
		}

		return &usingdb{
			GinHTTPHandler: httpHandler,
			userapiController: userapi.Controller{
				JWTToken:                  jwtToken,
				Env:                       env,
				Router:                    httpHandler.Router,
				CreateUserInport:          createuser.NewUsecase(datasource),
				ShowUserByIDInport:        showuserbyid.NewUsecase(datasource),
				ShowUserByEmailInport:     showuserbyemail.NewUsecase(datasource),
				ShowUserByUsernameInport:  showuserbyusername.NewUsecase(datasource),
				ShowAllUsersInport:        showallusers.NewUsecase(datasource),
				UpdateUserInport:          updateuser.NewUsecase(datasource),
				DeleteUserInport:          deleteuser.NewUsecase(datasource),
				ActivationUserInport:      activationuser.NewUsecase(datasource),
				ResetActivationUserInport: resetactivationuser.NewUsecase(datasource),
				LoginUserInport:           loginuser.NewUsecase(datasource),
			},
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *usingdb) SetupController() {
	r.userapiController.RegisterRouter()
	// TODO another router call will added here ... <<<<<<
}
