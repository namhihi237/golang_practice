package category

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Admin get categories
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "category id"
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/api/v1/categories [get]
func AdminGetCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		id := utils.StringToInt64(c.Param("id"))

		category, err := models.GetCategoryById(id)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, category)
	}
}
