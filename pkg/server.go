package pkg

import (
	"fmt"
	logutil "helpers/pkg/logutils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func RunServer(s IFileCollabServer) {
	logutil.Logger().Infof("Server is starting")
	err := s.InitializeDatabase()
	if err != nil {
		logutil.Logger().Errorf("Error while initializing database: %s. The server will stop", err.Error())
		return
	}
	// Set the Gin mode to release
	ginMode := os.Getenv("GIN_MODE")
	// Set the Gin mode based on the environment variable
	if ginMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		contextLogs, exists := params.Keys["contextLogs"]
		extraLogs := ""
		if exists {
			extraLogs = contextLogs.(string)
		}

		// Format the log output string
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\" %s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
			extraLogs,
		)
	}))
	base := engine.Group(s.GetBasePath())
	engine.MaxMultipartMemory = 8000 << 20
	engine.SetTrustedProxies(nil)
	s.LiveRoute(base)
	s.ReadyRoute(base)
	engine.Run(s.GetHost())
}
