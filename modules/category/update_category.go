package category

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"
	"practice/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryUpdate struct {
	Name        string `json:"name" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Description string `json:"description"`
}

// @Summary Update category
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param name body string true "name"
// @Param image body string true "image"
// @Param description body string false "description"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /categories/:id [PUT]
func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		validator := validator.New()

		id := utils.StringToInt64(c.Param("id"))

		var categoryInput CategoryUpdate
		if err := c.ShouldBind(&categoryInput); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		err := validator.Struct(categoryInput)
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, gin.H{"error": err.Error()})
			return
		}

		err = models.UpdateCategory(id, categoryInput.Name, categoryInput.Image, categoryInput.Description)

		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.INVALID_PARAMS, gin.H{"error": err.Error()})
			return
		}
		appG.Response(http.StatusOK, errors.SUCCESS, nil)
	}
}