package handler

import (
	"Golangcrud/pkg/service"

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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up",h.SignUp)
		auth.POST("/sign-in",h.SignIn)
	}

	api := router.Group("/api",h.userIdenity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/",h.createList)
			lists.GET("/",h.getAllList)
			lists.GET("/:id",h.getListById)
			lists.PUT("/:id",h.updateList)
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