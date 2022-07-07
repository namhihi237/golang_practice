package category

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Delete category
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "id category"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/api/v1/categories/:id [DELETE]
func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		id := utils.StringToInt64(c.Param("id"))
		if err := models.DeleteCategory(id); err != nil {
			appG.Response(http.StatusBadRequest, errors.NOT_FOUND, err.Error())
			return
		}
		appG.Response(http.StatusOK, errors.SUCCESS, nil)
	}
}
