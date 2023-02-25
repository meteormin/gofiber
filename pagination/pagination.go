package pagination

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Page struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`
}

type Paginator[T interface{}] struct {
	Page
	TotalCount int64 `json:"total_count"`
	Data       []T   `json:"data"`
}

func GetPageFromCtx(c *fiber.Ctx) (Page, error) {
	var page Page
	err := c.QueryParser(&page)
	if err != nil {
		return Page{Page: 1, PageSize: 10}, err
	}

	return page, nil
}

func Paginate(pageInfo Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pageInfo.Page
		if page <= 0 {
			page = 1
		}

		pageSize := pageInfo.PageSize
		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
