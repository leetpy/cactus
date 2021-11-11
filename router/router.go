package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/cactus/handler/project"
	"github.com/leetpy/cactus/handler/user"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	api := g.Group("/api")
	{
		api.GET("/project", project.List)
	}

	u := api.Group("/user")
	{
		u.POST("", user.Create)
		u.GET("/list", user.List)
		u.GET("/:username", user.Get)
	}
	return g
}
