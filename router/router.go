package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/goserver/handler/sd"
	"github.com/xiaozefeng/goserver/handler/user"
	"github.com/xiaozefeng/goserver/router/middleware"
	"net/http"
)

func Load(g *gin.Engine, mws ...gin.HandlerFunc) *gin.Engine {
	// add middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mws...)

	// add 404 router
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router")
	})

	v1u := g.Group("/v1/users")
	{
		v1u.POST("", user.Create)
		v1u.GET("", user.List)
		v1u.GET("/:username", user.Get)
		v1u.DELETE("/:id", user.Delete)
		v1u.PUT("/:id", user.Update)
	}

	// add health check handlers
	hg := g.Group("/sd")
	{
		hg.GET("/health", sd.HealthCheck)
		hg.GET("/disk", sd.DiskCheck)
		hg.GET("/cpu", sd.CPUCheck)
		hg.GET("/ram", sd.RAMCheck)
	}
	return g
}
