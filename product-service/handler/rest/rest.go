package restHandler

import (
	"go-service/usecase"
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

type rest struct {
	svc       *usecase.Usecase
	ginEngine *gin.Engine
	logger    *log.Logger
	opt       Option
}

type Option struct {
	JWTKey string `env:"JWT_KEY"`
	Port   string `env:"PORT"`
}

var once = &sync.Once{}

func Init(service *usecase.Usecase, ginEngine *gin.Engine, logger *log.Logger, opt Option) REST {
	var r *rest
	once.Do(func() {
		r = &rest{
			svc:       service,
			ginEngine: ginEngine,
			logger:    logger,
			opt:       opt,
		}
		r.Serve()
	})
	return r
}

func (r *rest) Serve() {

	r.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.ginEngine.Use(r.LoggerMiddleware())

	group := r.ginEngine.Group("/")
	{
		group.Use(r.AuthChecker())
		// Product
		group.GET("/product", r.GetProduct)
		group.POST("/product/:product_id/add-stock", r.AddStockProduct)
	}

}

func (r *rest) Run() {
	newGin := gin.New()
	newGin.Use(func(ctx *gin.Context) {
		r.ginEngine.ServeHTTP(ctx.Writer, ctx.Request)
	})
	port := ":" + r.opt.Port
	r.logger.Info("[HTTP] @", port)
	newGin.Run(port)
}
