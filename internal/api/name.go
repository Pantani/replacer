package api

import (
	"github.com/Pantani/logger"
	"github.com/Pantani/replacer/internal/regex"
	"github.com/Pantani/replacer/internal/storage"
	"github.com/Pantani/replacer/pkg/replacer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get Status
// @ID health
// @Description Get application status
// @Tags status
// @Success 200 {object} replacer.Status
// @Router /status [get]
func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, replacer.Status{Status: true})
}

// @Summary Insert name
// @ID insert_name
// @Description Insert or update a name
// @Accept json
// @Produce json
// @Tags name
// @Param name path string true "the name"
// @Param url body replacer.NameRequest true "The url to insert"
// @Success 200 {object} replacer.Status
// @Error 404 {object} replacer.Error
// @Router /names/{name} [put]
func insertName(storage storage.Name) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		name := c.Param("name")
		if !regex.IsValidName(name) {
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: "invalid name"})
			return
		}

		var req replacer.Name
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: "invalid body"})
			return
		}
		req.Name = name
		err := storage.InsertOrUpdateName(req)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, replacer.Status{Status: true})
	}
}

// @Summary Get name
// @ID get_name
// @Description Get a name detail
// @Accept json
// @Produce json
// @Tags name
// @Param name path string true "the name"
// @Success 200 {object} replacer.Status
// @Error 404 {object} replacer.Error
// @Router /names/{name} [get]
func getName(storage storage.Name) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		name := c.Param("name")
		if !regex.IsValidName(name) {
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: "invalid name"})
			return
		}
		r, err := storage.GetName(name)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

// @Summary Get all names
// @ID get_all_names
// @Description Get all names details
// @Accept json
// @Produce json
// @Tags name
// @Success 200 {object} replacer.Status
// @Error 404 {object} replacer.Error
// @Router /names [get]
func getAllNames(storage storage.Name) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		r, err := storage.GetAllNames()
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

// @Summary Delete name
// @ID del_name
// @Description Delete a name
// @Accept json
// @Produce json
// @Tags name
// @Param name path string true "the name"
// @Success 200 {object} replacer.Status
// @Error 404 {object} replacer.Error
// @Router /names/{name} [delete]
func deleteName(storage storage.Name) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		name := c.Param("name")
		if !regex.IsValidName(name) {
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: "invalid name"})
			return
		}
		err := storage.DeleteName(name)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, replacer.Status{Status: true})
	}
}

// @Summary Delete all names
// @ID del_all_names
// @Description Delete all name
// @Accept json
// @Produce json
// @Tags name
// @Success 200 {object} replacer.Status
// @Error 404 {object} replacer.Error
// @Router /names [delete]
func deleteAllNames(storage storage.Name) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		err := storage.DeleteAllNames()
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, replacer.Status{Status: true})
	}
}
