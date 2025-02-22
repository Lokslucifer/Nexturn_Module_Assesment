package controller

import (
	"blog_server/model"
	"blog_server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogservice *service.BlogService) *BlogController {
	return &BlogController{BlogService: blogservice}
}

func (controller *BlogController) CreateBlog(c *gin.Context) {
	var blog model.Blog
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdBlog)

}

func (controller *BlogController) GetBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	blog, err := controller.BlogService.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)

}

func (controller *BlogController) GetAllBlogs(c *gin.Context) {
	blogs, err := controller.BlogService.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (controller *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var blog model.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.Id = blogID
	updatedBlog, err := controller.BlogService.UpdateBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

func (controller *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = controller.BlogService.DeleteBlog(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
