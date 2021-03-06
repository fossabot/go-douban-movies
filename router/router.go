package router

import (
	"github.com/gin-gonic/gin"

	"github.com/linehk/go-douban-movies/router/api/v1"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/movies", v1.Movies)
	}
	return r
}
