package product

import (
	"net/http"
	"practice/models"
	"practice/pkg/app"
	"practice/pkg/errors"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" validate:"required"`
	Amount      int64    `json:"amount" default:"0"`
	CategoryIds []int64  `json:"category_ids" required:"required"`
	ImagesUrls  []string `json:"image_urls" validate:"required"`
}

// @Summary Add product
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param name body string true "name"
// @Param image body string true "image"
// @Param description body string false "description"
// @Param price body number true "price"
// @Param amount body int true "amount"
// @Param category_ids body []int true "category_ids"
// @Param image_urls body []string true "image urls"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/api/v1/products [post]
func AdminAddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		validator := validator.New()

		var productInput ProductInput
		if err := c.ShouldBind(&productInput); err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, nil)
			return
		}

		err := validator.Struct(productInput)
		if err != nil {
			appG.Response(http.StatusBadRequest, errors.INVALID_PARAMS, gin.H{"error": err.Error()})
			return
		}

		validateCategoryIds, err := models.CheckCategoryIds(productInput.CategoryIds)
		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.INVALID_PARAMS, gin.H{"error": err.Error()})
			return
		}
		if !validateCategoryIds {
			appG.Response(http.StatusInternalServerError, errors.CATEGORY_IDS_INVALID_PARAMS, nil)
			return
		}

		product := models.Product{
			Name:        productInput.Name,
			Description: productInput.Description,
			Price:       productInput.Price,
			Amount:      productInput.Amount,
		}

		err = models.CreateProduct(productInput.CategoryIds, productInput.ImagesUrls, &product)

		if err != nil {
			appG.Response(http.StatusInternalServerError, errors.INVALID_PARAMS, gin.H{"error": err.Error()})
			return
		}

		appG.Response(http.StatusOK, errors.SUCCESS, nil)

	}
}
