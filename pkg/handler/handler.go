package handler

import (
	"github.com/gin-gonic/gin"

	_ "github.com/Murolando/hakaton_geo/docs"
	"github.com/Murolando/hakaton_geo/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	// router.Static("/storage","./storage")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.GET("/refresh/:refresh", h.newRefresh)
		}
		class := api.Group("/class")
		{
			class.GET("/dashboard", h.userIdentity, h.dashboardClass)
			class.GET("/my-classes", h.userIdentity, h.myClass)
			class.GET("/:class_id", h.userIdentity, h.oneClass)
			// class.DELETE("/class_id",)
			// class.PUT("/class_id",)
			// class.POST("/",)

		}
		kontur := api.Group("/kontur")
		{
			kontur.GET("/start-game/:n", h.startKonturGame)
			kontur.PUT("/process-game", h.userIdentity, h.processKonturGame)
		}
	}
	return router
}
