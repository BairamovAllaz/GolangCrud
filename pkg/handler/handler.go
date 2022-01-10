package handler

import (
	"Golangcrud/pkg/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) Initroutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up",h.SignUp)
		auth.POST("/sign-in",h.SignIn)
		auth.POST("/forgot-password",h.ForgotPassword)
		auth.POST("/forgot-password/:token",h.ForgotPasswordHandler)
	}

	api := router.Group("/api",h.userIdenity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/",h.createList)
			lists.GET("/",h.getAllList)
			lists.GET("/:id",h.getListById)
			lists.PATCH("/:id",h.updateList)
			lists.DELETE("/:id",h.deleteList)
		}

		items := lists.Group(":id/items")
		{
			items.POST("/",h.createItem)
			items.GET("/",h.getAllItems)
			items.GET("/:item_id",h.getItemById)
			items.PUT("/:item_id",h.updateItem)
			items.DELETE("/:item_id",h.deleteItem)
		}

	}
	return router
}
