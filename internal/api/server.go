package api

import (
	"github.com/gin-gonic/gin"
	server "helpers/pkg"
	"net/http"
)

type FileCollabServer struct {
	server.FileCollabServer
}

func (s FileCollabServer) GetBasePath() string {
	return "/files"
}

func (s FileCollabServer) RouteReady(router gin.IRouter) {
	router.GET("/ready", func(ctx *gin.Context) {

		ctx.JSON(200, http.DefaultClient)
	})
}
func (s FileCollabServer) FunctionnalRoutes(router gin.IRouter) {

	router.GET("/ff", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "dodo"})
	})
}

func (s FileCollabServer) RouteLive(router gin.IRouter) {
	router.GET("/live", func(ctx *gin.Context) {
		ctx.JSON(200, http.DefaultClient)
	})

}
