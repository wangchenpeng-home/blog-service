package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/wangchenpeng-home/blog-service/docs"
	"github.com/wangchenpeng-home/blog-service/global"
	"github.com/wangchenpeng-home/blog-service/internal/middleware"
	v1 "github.com/wangchenpeng-home/blog-service/internal/routers/api/v1"
	"github.com/wangchenpeng-home/blog-service/pkg/limiter"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInertval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Tracing())
	//注册翻译中间件
	r.Use(middleware.Translations())
	//增加接口文档查看
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := v1.NewUpload()

	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UpLoadSavePath))
	r.GET("/auth", v1.GetAuth)
	apiV1 := r.Group("/api/v1")
	// 注册JWT加密中间件
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags:id", tag.Delete)
		apiV1.PUT("/tags:id", tag.Update)
		apiV1.PATCH("/tags:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles:id", article.Delete)
		apiV1.PUT("/articles:id", article.Update)
		apiV1.PATCH("/articles:id/state", article.Update)
		apiV1.GET("/articles:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
