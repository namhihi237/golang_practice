package utils

import (
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int   `form:"page" binding:"required,numeric" default:"1" json:"page"`
	Limit int   `form:"limit" binding:"required,numeric" default:"10" json:"limit"`
	Total int64 `form:"-" json:"total"`
}

func GetPaging(c *gin.Context) (*Pagination, error) {
	var paging Pagination
	if err := c.ShouldBindQuery(&paging); err != nil {
		return nil, err
	}
	return &paging, nil
}

func (p *Pagination) DefaultPaging() *Pagination {
	return &Pagination{
		Page:  1,
		Limit: 10,
	}
}
