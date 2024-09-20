package pkg

import (
	"github.com/gin-gonic/gin"
	"helpers/sharedconfigs"
)

type IFileCollabServer interface {
	InitializeDatabase() error
	GetHost() string
	LiveRoute(router gin.IRouter)
	ReadyRoute(router gin.IRouter)
	GetBasePath() string
}

type FileCollabServer struct {
	ServerConfig sharedconfigs.IServerConfiguration
}

func (s FileCollabServer) InitializeDatabase() error {
	return nil
}

func (s FileCollabServer) GetHost() string {
	return s.ServerConfig.GetHost()
}

func (s FileCollabServer) GetBasePath() string {
	return s.ServerConfig.GetBasePath()
}

func (s FileCollabServer) LiveRoute(router gin.IRouter) {
	/* monitoringMiddlewre := auth_v2.GetOneCheckerMiddleware(permission.Monitor, true)
	router.GET("/live", monitoringMiddlewre, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	},
	) */
	router.GET("/live", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	},
	)
}

func (s FileCollabServer) ReadyRoute(router gin.IRouter) {
	/* monitoringMiddlewre := auth_v2.GetOneCheckerMiddleware(permission.Monitor, true)
	router.GET("/ready", monitoringMiddlewre, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	},
	) */
	router.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	},
	)
}
