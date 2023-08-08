package dto

// Ini buat paging di taruh di parameter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

// ini buat paging di taruh di return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

// ini buat di taruh di response
type Paging struct {
	Page        int `json:"paging"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}
