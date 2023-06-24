package handler

import (
	"github.com/OOO-Roadmap-Creation/roadmap2/pkg/repository"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/OOO-Roadmap-Creation/roadmap2/docs"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(services *repository.Repository) *Handler {
	return &Handler{repo: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	publishedRoadmap := router.Group("/published-roadmap")
	{
		publishedRoadmap.GET("/", h.list)
		publishedRoadmap.PUT("/visibility/:id/:v", h.visibility)
		publishedRoadmap.GET("/:p1", h.common)
		publishedRoadmap.GET("/:p1/:p2", h.common)
	}
	publishedNode := router.Group("/published-node")
	{
		publishedNode.GET("/roadmap/:id", h.nodes)
	}
	rating := router.Group("/rating/roadmap")
	{
		rating.GET("/:id", h.getRating)
		rating.GET("/:id/user", h.getRatingUser)
		rating.POST("/:id", h.setRating)
	}

	return router
}
