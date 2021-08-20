package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GinHTTPHandler will define basic HTTP configuration with gracefully shutdown
type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
}

func NewGinHTTPHandler(address string) (GinHTTPHandler, error) {

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ready")
	})

	return GinHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(router, address),
		Router:             router,
	}, nil

}

// RunApplication is implementation of RegistryContract.RunApplication()
func (r *GinHTTPHandler) RunApplication() {
	r.RunWithGracefullyShutdown()
}
