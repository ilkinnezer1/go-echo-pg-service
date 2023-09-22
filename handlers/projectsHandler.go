package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"io"
	cacheInit "main/cache"
	"main/models"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetSingleProjectDetails(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid partner ID",
		})
	}

	cacheKey := "project" + strconv.Itoa(id)
	if project, found := cacheInit.CacheInstance.Get(cacheKey); found {
		return c.JSON(http.StatusOK, project)
	}

	project := new(models.Projects)
	if err := db.First(project, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Project not found",
		})
	}
	project.ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(project.ImagePath)

	// Add the single project to the cache
	cacheInit.CacheInstance.Set(cacheKey, project, cache.DefaultExpiration)
	cacheInit.CacheInstance.Delete(cacheKey)
	return c.JSON(http.StatusOK, project)
}

func GetProjectsDetails(c echo.Context) error {
	cacheKey := "projects"
	if partners, found := cacheInit.CacheInstance.Get(cacheKey); found {
		return c.JSON(http.StatusOK, partners)
	}

	var projectsDetails []models.Projects
	// Return Projects Details with asc order
	if err := db.Order("id asc").Find(&projectsDetails).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	// Map the projects and include to the full file URL
	for i := range projectsDetails {
		projectsDetails[i].ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(projectsDetails[i].ImagePath)
	}
	// Set the partner to the cache
	cacheInit.CacheInstance.Set(cacheKey, projectsDetails, cache.DefaultExpiration)

	return c.JSON(http.StatusOK, projectsDetails)
}

func CreateNewProject(c echo.Context) error {
	// Create New Instance of Project
	projects := new(models.Projects)

	if err := c.Bind(projects); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	// Get the uploaded file(image)
	uploadedFile, err := c.FormFile("image")

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	src, err := uploadedFile.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	defer func(src multipart.File) {
		var err = src.Close()
		if err != nil {

		}
	}(src)

	// Replaces spaces with underscores on file name
	filename := strings.Replace(uploadedFile.Filename, " ", "_", -1)

	// Full destination of the uploaded file
	uploadedFileDst := filepath.Join("uploads/projects", filename)
	dst, err := os.Create(uploadedFileDst)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	defer func(dst *os.File) {
		var err = dst.Close()
		if err != nil {

		}
	}(dst)

	// Copy the uploaded file dest
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Attach the full file destination to the model property
	projects.ImagePath = uploadedFileDst
	projects.Name = c.FormValue("name")
	projects.Slogan = c.FormValue("slogan")
	projects.ImgAltText = c.FormValue("imgAltText")
	projects.ShortIntro = c.FormValue("shortIntro")
	projects.Hyperlink = c.FormValue("hyperlink")

	if err := db.Create(&projects).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	projects.ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(projects.ImagePath)

	cacheInit.CacheInstance.Delete("projects")
	// Add the new partner to the cache
	cacheInit.CacheInstance.Set("project: "+strconv.Itoa(projects.ID), projects, cache.DefaultExpiration)

	return c.JSON(http.StatusCreated, projects)
}

func UpdateProjectDetail(c echo.Context) error {
	projectDetail := new(models.Projects)

	if err := c.Bind(projectDetail); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	id, err := strconv.Atoi(c.Param("id"))

	bannerImg, err := c.FormFile("image")

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := db.First(&projectDetail, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Delete existing image from uploads via id
	if err := os.Remove(projectDetail.ImagePath); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	bannerImgName := strings.Replace(bannerImg.Filename, " ", "_", -1)

	src, err := bannerImg.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	defer func(src multipart.File) {
		var err = src.Close()
		if err != nil {

		}
	}(src)

	fileDestination := filepath.Join("uploads/projects", bannerImgName)
	destination, err := os.Create(fileDestination)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	defer func(destination multipart.File) {
		var err = destination.Close()
		if err != nil {
		}
	}(destination)

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(destination, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	projectDetail.Name = c.FormValue("name")
	projectDetail.Slogan = c.FormValue("slogan")
	projectDetail.ShortIntro = c.FormValue("shortIntro")
	projectDetail.ImgAltText = c.FormValue("imgAltText")
	projectDetail.Hyperlink = c.FormValue("hyperlink")
	projectDetail.ImagePath = fileDestination

	if err := db.Save(&projectDetail).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	projectDetail.ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(projectDetail.ImagePath)

	updatedProject := models.Projects{
		ID:         id,
		Name:       projectDetail.Name,
		Slogan:     projectDetail.Slogan,
		ShortIntro: projectDetail.ShortIntro,
		ImagePath:  projectDetail.ImagePath,
		ImgAltText: projectDetail.ImgAltText,
		Hyperlink:  projectDetail.Hyperlink,
	}

	cacheInit.CacheInstance.Delete("projects")
	cacheInit.CacheInstance.Set("project: "+strconv.Itoa(id), updatedProject, cache.DefaultExpiration)

	return c.JSON(http.StatusOK, projectDetail)
}

func DeleteProject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := db.Where("id = ?", id).Delete(&models.Projects{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	cacheInit.CacheInstance.Delete("project:" + strconv.Itoa(id))
	cacheInit.CacheInstance.Delete("projects")

	return c.JSON(http.StatusNoContent, map[string]string{
		"message": "Project Details Successfully removed",
	})
}
