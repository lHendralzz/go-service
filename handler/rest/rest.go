package restHandler

import (
	"go-service/service"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type REST interface {
	Serve()
	Run()
}

type rest struct{
    svc *service.Service
    ginEngine *gin.Engine
    logger *log.Logger
}

var once = &sync.Once{}

func Init(service *service.Service, ginEngine *gin.Engine, logger *log.Logger) REST {
    var r *rest
    once.Do(func(){
        r = &rest{
            svc: service,
            ginEngine: ginEngine,
            logger: logger,
        }
        r.Serve()
    })
    return r
}

func (r *rest) Serve() {

    r.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.ginEngine.Use(r.LoggerMiddleware())
    r.ginEngine.Handle(POST, "/login", r.Login)
    group := r.ginEngine.Group("")
    {
        group.Use(r.AuthChecker()) 
        group.Handle(GET, "/testing2", r.Testing)
    }
}

func (r *rest) Run() {
    newGin := gin.New()
    newGin.Use(func (ctx *gin.Context){
        r.ginEngine.ServeHTTP(ctx.Writer, ctx.Request)
    })
    port := ":8080";
    r.logger.Info("[HTTP] @", port)
    newGin.Run(port)
}