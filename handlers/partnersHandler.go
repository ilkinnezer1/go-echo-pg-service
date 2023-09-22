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

func GetSinglePartner(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid partner ID",
		})
	}

	cacheKey := "partner:" + strconv.Itoa(id)
	if partner, found := cacheInit.CacheInstance.Get(cacheKey); found {
		return c.JSON(http.StatusOK, partner)
	}

	partner := new(models.Partners)
	if err := db.First(partner, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Partner not found",
		})
	}

	partner.ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(partner.ImagePath)

	cacheInit.CacheInstance.Set(cacheKey, partner, cache.DefaultExpiration)
	cacheInit.CacheInstance.Delete(cacheKey)
	return c.JSON(http.StatusOK, partner)
}

func GetPartners(c echo.Context) error {
	cacheKey := "partners"
	if partners, found := cacheInit.CacheInstance.Get(cacheKey); found {
		return c.JSON(http.StatusOK, partners)
	}
	var partners []models.Partners
	// Return partners logo with asc order
	if err := db.Order("id asc").Find(&partners).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	// Map the partners to include the full image URL
	for i := range partners {
		partners[i].ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(partners[i].ImagePath)
	}

	cacheInit.CacheInstance.Set(cacheKey, partners, cache.DefaultExpiration)
	return c.JSON(http.StatusOK, partners)
}

func CreatePartner(c echo.Context) error {
	// Create new instance of partners
	partner := new(models.Partners)
	if err := c.Bind(partner); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	// Get the uploaded file
	partner.Title = c.FormValue("title")
	partner.AltText = c.FormValue("altText")
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Replaces spaces with underscores on file name
	filename := strings.Replace(file.Filename, " ", "_", -1)

	// Open the file
	src, err := file.Open()
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

	// Get the destination of the uploaded file
	fileDst := filepath.Join("uploads/partners", filename)
	dst, err := os.Create(fileDst)

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

	// Copy the uploaded file to the dest file
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	// Attach to the model
	partner.ImagePath = fileDst

	if err := db.Create(&partner).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	partner.ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(partner.ImagePath)

	cacheInit.CacheInstance.Delete("partners")
	cacheInit.CacheInstance.Set("partner: "+strconv.Itoa(partner.ID), partner, cache.DefaultExpiration)

	return c.JSON(http.StatusCreated, partner)
}

func UpdatePartner(c echo.Context) error {
	partner := new(models.Partners)
	if err := c.Bind(partner); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := db.First(&partner, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Partner not found",
		})
	}

	// Delete existing partner image via id
	if err := os.Remove(partner.ImagePath); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	filename := strings.Replace(file.Filename, " ", "_", -1)

	src, err := file.Open()
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

	fileDst := filepath.Join("uploads/partners", filename)
	dst, err := os.Create(fileDst)
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

	// Copy the uploaded file to the dest file
	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	partner.Title = c.FormValue("title")
	partner.AltText = c.FormValue("altText")
	partner.ImagePath = fileDst

	if err := db.Save(&partner).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	partner.ImagePath = c.Scheme() + "://" + c.Request().Host + "/" + filepath.ToSlash(partner.ImagePath)

	updatedPartner := models.Partners{
		ID:        id,
		Title:     partner.Title,
		ImagePath: partner.ImagePath,
		AltText:   partner.AltText,
	}

	cacheInit.CacheInstance.Delete("partners")
	cacheInit.CacheInstance.Set("partner: "+strconv.Itoa(id), updatedPartner, cache.DefaultExpiration)

	return c.JSON(http.StatusOK, partner)
}

func DeletePartner(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := db.Where("id = ?", id).Delete(&models.Partners{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	cacheInit.CacheInstance.Delete("partner:" + strconv.Itoa(id))
	cacheInit.CacheInstance.Delete("partners")

	return c.JSON(http.StatusNoContent, map[string]string{
		"message": "Content is successfully removed",
	})
}
