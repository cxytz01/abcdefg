package busi

import (
	"fmt"
	v1 "producer/internal/busi/api/v1"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setConfig(csvstore string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.CSVStore, csvstore)

		c.Next()
	}
}

func (s *HttpServer) registerV1(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)

		apiv1.POST("/campaign", setConfig(s.csvstore), v1.CreateCampaign) // -Call by client, Create campaingn
		apiv1.POST("/messages", v1.DispatchToKafka)                                // -Call by scheduler

	}
}

func (s *HttpServer) RegisterRoutes(r *gin.Engine) {
	// r.Use(utils.Cors())
	r.Use(cors.Default())
	r.GET("/producer/swagger/*any", swagHandler)

	s.registerV1(r)
}

func (s *HttpServer) Start() {
	// // if Flags.Mode == "prod" {
	gin.SetMode(gin.ReleaseMode)
	// // }

	// r := gin.Default()
	r := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())
	s.RegisterRoutes(r)
	r.Run(s.addr)
}

type HttpServer struct {
	addr     string
	csvstore string
}

func NewHttpServer(addr, csvstore string) *HttpServer {
	return &HttpServer{addr, csvstore}
}

func HttpServerStart(addr, csvstore string) {
	NewHttpServer(addr, csvstore).Start()
}
