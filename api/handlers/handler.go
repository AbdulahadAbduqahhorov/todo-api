package handlers

import (
	"github.com/AbdulahadAbduqahhorov/gin/todo-api/api/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}
func (h *Handler) InitializeHandler() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}

	api := router.Group("/api",h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			todos := lists.Group(":id/todo")
			{
				todos.POST("/", h.createItem)
				todos.GET("/", h.getAllItems)
				todos.GET("/:todo_id", h.getItemByID)
				todos.PUT("/:todo_id", h.updateItem)
				todos.DELETE("/:todo_id", h.deleteItem)
			}
		}

	}
	return router
}
