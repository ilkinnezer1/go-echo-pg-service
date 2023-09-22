package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	cacheInit "main/cache"
	"main/models"
	"net/http"
	"strconv"
)

func GetSingleBlog(c echo.Context) error {
	// Parse the id from the request URL
	id := c.Param("id")
	cacheKey := "blog" + id

	if blog, found := cacheInit.CacheInstance.Get(cacheKey); found {
		return c.JSON(http.StatusOK, blog)
	}

	var blog models.Blog
	if err := db.First(&blog, id).Error; err != nil {
		// Return an error response if the blog is not found
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Blog not found",
		})
	}

	// Add the blog post to the cache
	cacheInit.CacheInstance.Set(cacheKey, blog, cache.DefaultExpiration)
	cacheInit.CacheInstance.Delete(cacheKey)

	// Return the blog post as a JSON response
	return c.JSON(http.StatusOK, blog)
}

func GetAllBlogs(c echo.Context) error {
	// Set the new cache entry
	cacheKey := "blogs"
	if blogs, found := cacheInit.CacheInstance.Get(cacheKey); found {
		return c.JSON(http.StatusOK, blogs)
	}

	var blogs []models.Blog
	if err := db.Order("pinned desc, COALESCE(updated_at, created_at) desc").Find(&blogs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	// Set the blog to the cache
	cacheInit.CacheInstance.Set(cacheKey, blogs, cache.DefaultExpiration)

	return c.JSON(http.StatusOK, blogs)
}

func CreateBlog(c echo.Context) error {
	// Create new instance of Blog model
	blog := new(models.Blog)

	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := db.Create(&blog).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	cacheInit.CacheInstance.Delete("blogs")
	// Add the new blog to the cache
	cacheInit.CacheInstance.Set("blog:"+strconv.Itoa(blog.ID), blog, cache.DefaultExpiration)
	// Return created blog as a JSON response
	return c.JSON(http.StatusCreated, blog)
}

func UpdateBlog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	// Create a new instance of Blog model
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := db.Model(&models.Blog{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       blog.Title,
		"subtitle":    blog.Subtitle,
		"description": blog.Description,
		"hyperlink":   blog.Hyperlink,
		"pinned":      blog.Pinned,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	// Add the updated blog to the cache
	updatedBlog := models.Blog{
		ID:          id,
		Title:       blog.Title,
		Subtitle:    blog.Subtitle,
		Description: blog.Description,
		Hyperlink:   blog.Hyperlink,
		Pinned:      blog.Pinned,
	}
	cacheInit.CacheInstance.Set("blog:"+strconv.Itoa(id), updatedBlog, cache.DefaultExpiration)

	// Delete the cache entry for all blogs
	cacheInit.CacheInstance.Delete("blogs")
	return c.JSON(http.StatusOK, blog)
}

func DeleteBlog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := db.Delete(&models.Blog{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	// Delete the cache entry for this blog
	cacheInit.CacheInstance.Delete("blog:" + strconv.Itoa(id))
	cacheInit.CacheInstance.Delete("blogs")

	return c.NoContent(http.StatusNoContent)
}
