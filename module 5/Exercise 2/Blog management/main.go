package main

import (
	db "blog_server/config"
	"blog_server/controller"
	"blog_server/middleware"
	"blog_server/repository"
	"blog_server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitialiseDatabase()
	userRepo := repository.NewBlogRepository(db.GetDB())
	userService := service.NewBlogService(userRepo)
	userController := controller.NewBlogController(userService)

	// Initialize Gin router
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware())
	protected := r.Group("/blogs", middleware.AuthMiddleware())
	{
		protected.POST("", userController.CreateBlog)
		protected.PUT("/:id", userController.UpdateBlog)
		protected.DELETE("/:id", userController.DeleteBlog)
	}

	r.GET("/blogs/:id", userController.GetBlog) // Public
	r.GET("/blogs", userController.GetAllBlogs) // Public

	// Routes
	r.POST("/blogs", userController.CreateBlog)
	r.GET("/blogs/:id", userController.GetBlog)
	r.GET("/blogs", userController.GetAllBlogs)
	r.PUT("/blogs/:id", userController.UpdateBlog)
	r.DELETE("/blogs/:id", userController.DeleteBlog)

	// Start server
	r.Run(":8000")

}
