package utilQuery

import (
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

func Pagination(query *gorm.DB, queryValues map[string][]string) *gorm.DB {
	q := url.Values(queryValues)
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query = query.Offset(offset).Limit(pageSize) // Pagination

	return query

}
