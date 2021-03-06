package category

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
)

// note: https://gorm.io/docs/advanced_query.html#Smart-Select-Fields

// @Summary Admin get categories
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/api/v1/categories [get]
func AdminGetListCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		paging, err := utils.GetPaging(c)
		if err != nil {
			paging = paging.DefaultPaging()
		}

		categories, err := models.GetCategories(paging.Page, paging.Limit)
		paging.Total = models.CountCategory()

		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.SERVER_ERROR, nil)
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, map[string]interface{}{
			"categories": categories,
			"paging":     paging,
		})
	}
}
