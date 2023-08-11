package common

import (
	"employeeleave/model/dto"
	"math"
	"os"
	"strconv"
)

func GetPaginationParams(params dto.PaginationParam) dto.PaginationQuery {
	// err := LoadEnv()
	// execptions.CheckErr(err)

	var (
		page, take, skip int
	)

	if params.Page > 0 {
		page = params.Page
	} else {
		page = 1
	}

	if params.Limit == 0 {
		n, _ := strconv.Atoi(os.Getenv("DEFAULT_ROWS_PER_PAGE"))
		take = n
	} else {
		take = params.Limit
	}

	// rumus offset / rumus pagination
	// product => 10 | Page 1 => row 1 s.d 5
	// product => 10 | Page 2 => row 6 s.d 10
	// offset = (page - 1) * limit
	// offset = (1 - 1) * 5 ==== 0
	// offset = (2 - 1) * 5 ==== 5
	// SELECT * FROM product LIMIT 5 OFFSET 5

	if page > 0 {
		skip = (page - 1) * take
	} else {
		skip = 0
	}

	return dto.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

// 21 / 5 === 4.xxx
// ceil (pembulatan keatas) e.g 4.2 == 4 | 4.6 == 5
func Paginate(page, limit, totalRows int) dto.Paging {
	return dto.Paging{
		Page:        page,
		RowsPerPage: limit,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}
