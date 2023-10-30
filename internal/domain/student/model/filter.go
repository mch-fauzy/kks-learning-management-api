package model

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"

	OperatorEqual  = "eq"
	OperatorRange  = "range"
	OperatorIn     = "in"
	OperatorIsNull = "is_null"
	OperatorNot    = "not"
)

type Filter struct {
	SelectFields []string      `json:"fields"`
	Sorts        []Sort        `json:"sort"`
	Pagination   Pagination    `json:"pagination"`
	FilterFields []FilterField `json:"filter"`
}

type Sort struct {
	Field string `json:"field"`
	Order string `json:"order" validate:"oneof=ASC DESC"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type FilterField struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator" validate:"oneof=eq range in is_null not"`
	Value    interface{} `json:"value"`
}
