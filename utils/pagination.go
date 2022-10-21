package utils

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	SortValue  string `json:"sort_value,omitempty;query:sortvalue"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 5
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at "
	}
	return p.Sort
}
func (p *Pagination) GetSortValue() string {
	if p.SortValue == "" {
		p.SortValue = "asc"
	}
	return p.SortValue
}

func (p *Pagination) GetPagination(c *gin.Context) Pagination {
	limit := p.GetLimit()
	page := p.GetPage()
	sort := p.GetSort()
	sortValue := p.GetSortValue()

	query := c.Request.URL.Query()
	for key, val := range query {
		queryValue := val[len(val)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "sort_value":
			sortValue = queryValue
		}
	}

	return Pagination{Limit: limit, Page: page, Sort: sort, SortValue: sortValue}
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort() + " " + pagination.GetSortValue())
	}
}
