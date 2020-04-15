package api

import (
	"github.com/Pantani/logger"
	"github.com/Pantani/replacer/internal/annotate"
	"github.com/Pantani/replacer/internal/storage"
	"github.com/Pantani/replacer/pkg/replacer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get Annotate
// @ID get_annotate
// @Description Get all annotates html snippet
// @Accept plain
// @Produce plain
// @Tags annotate
// @Param html body string true "The html to be replaced"
// @Success 200 {object} replacer.Html
// @Error 404 {object} replacer.Error
// @Router /annotate [post]
func generateAnnotate(storage storage.Name) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		names, err := storage.GetAllNames()
		if err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: err.Error()})
			return
		}
		data, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, replacer.Error{Message: "invalid body"})
			return
		}
		r := annotate.ScrapeHtml(string(data), names)
		c.String(http.StatusOK, r)
	}
}
