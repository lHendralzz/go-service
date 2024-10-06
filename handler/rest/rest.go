package restHandler

import (
	"go-service/service"
	"sync"

	"github.com/gin-gonic/gin"
)

type REST interface {
	Serve()
	Run()
}

type rest struct{
    svc *service.Service
    ginEngine *gin.Engine
}

var once = &sync.Once{}

func Init(service *service.Service,ginEngine *gin.Engine) REST {
    var r *rest
    once.Do(func(){
        r = &rest{
            svc: service,
            ginEngine: ginEngine,
        }
        r.Serve()
    })
    return r
}

func (r *rest) Serve() {
    r.ginEngine.Handle(GET, "/testing", r.Testing)
}

func (r *rest) Run() {
    newGin := gin.New()
    newGin.Use(func (ctx *gin.Context){
        r.ginEngine.ServeHTTP(ctx.Writer, ctx.Request)
    })
    port := ":8080";
    newGin.Run(port)
}
