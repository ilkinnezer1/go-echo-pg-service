package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/admin"
	"main/handlers"
	"main/token"
)

func BaseRouter() *echo.Echo {
	e := echo.New()

	// Middleware For Logs
	e.Use(handlers.LoggerMiddlewareConfig)
	e.Use(middleware.BodyDump(handlers.HTTPBodyDumpResponse))
	e.Use(middleware.Recover())

	// Solves the CORS error for the specific URL
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:1234"},
		AllowMethods: []string{echo.GET, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	// Auth Router
	e.POST("/login", handlers.LoginHandler)

	// Routers w/o protection
	e.GET("/blogs", handlers.GetAllBlogs)
	e.GET("/blogs/:id", handlers.GetSingleBlog)

	e.GET("/projects", handlers.GetProjectsDetails)
	e.GET("/projects/:id", handlers.GetSingleProjectDetails)

	e.GET("/partners", handlers.GetPartners)
	e.GET("/partners/:id", handlers.GetSinglePartner)

	// Protected Group URLs
	blogs := e.Group("/blog")
	partner := e.Group("/partner")
	projects := e.Group("/project")

	// Middleware to authenticate admin users
	blogs.Use(admin.AuthMiddleware)
	blogs.Use(echojwt.WithConfig(token.JwtConfig))

	blogs.POST("/create", handlers.CreateBlog)
	blogs.PATCH("/:id", handlers.UpdateBlog)
	blogs.DELETE("/:id", handlers.DeleteBlog)

	// Auth Middleware for Partner section
	partner.Use(admin.AuthMiddleware)
	partner.Use(echojwt.WithConfig(token.JwtConfig))

	// Partner section group
	partner.POST("/create", handlers.CreatePartner)
	partner.PATCH("/:id", handlers.UpdatePartner)
	partner.DELETE("/:id", handlers.DeletePartner)

	// Auth Middleware for Project Section
	projects.Use(admin.AuthMiddleware)
	projects.Use(echojwt.WithConfig(token.JwtConfig))

	// Project Section group
	projects.POST("/create", handlers.CreateNewProject)
	projects.PATCH("/:id", handlers.UpdateProjectDetail)
	projects.DELETE("/:id", handlers.DeleteProject)

	return e
}
