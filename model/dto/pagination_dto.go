package dto

// ini buat paging ditaruh di parameter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

// ini buat paging ditaruh di parameter
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

// ini buat paging ditaruh di response
type Paging struct {
	Page        int
	RowsPerPage int
	TotalRows   int
	TotalPages  int
}

// product 100
// paging {page: 1, RowsPerPage: 10, TotalRows: 100, TotalPages: 10}
