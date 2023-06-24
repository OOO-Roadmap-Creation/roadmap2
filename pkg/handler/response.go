package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type publishedRoadmapResponse struct {
	Id            interface{} `json:"id"`
	Version       int         `json:"version"`
	Visible       bool        `json:"visible"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	DateOfPublish time.Time   `json:"dateOfPublish"`
}

type publishedNodeResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	ParentId    int    `json:"parentId"`
}

type listRoadmapResponse struct {
	key []publishedRoadmapResponse
}
