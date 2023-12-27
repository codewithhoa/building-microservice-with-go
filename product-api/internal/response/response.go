package response

type MetaData struct {
	// The PageSize is total items of one page
	//
	// min: 1
	// Example: 10
	PageSize int `json:"pageSize"`

	// The Page is page which user want retrieve
	//
	// min: 1
	// Example: 2
	Page int `json:"page"`

	// The TotalPages is total of page
	//
	// Example: 100
	TotalPages int `json:"totalPages"`

	// The TotalRecords is total records
	//
	// Example: 1000
	TotalRecords int `json:"totalRecords"`

	// The sort
	//
	// Example: ["+createdAt", "-name"]
	Sort []string `json:"sort"`

	// The HasNext is true if have result in the next page
	//
	// Example: true
	HasNext bool `json:"hasNext"`
}

type SuccessMultiResponse struct {
	Data     any      `json:"data"`
	Metadata MetaData `json:"metadata"`
}

type SuccessSingleResponse struct {
	Data any `json:"data"`
}

type ErrorField struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

type ErrorResponse struct {
	// The Error is general error
	//
	// Example: "field1: required, field2: greater than 10"
	Error string `json:"error"`

	// The Code is code name of the error
	//
	// Example: "CREATE_BODY_INVALID"
	Code string `json:"code"`

	// The Errors is list of errors
	//
	// Example: [{"field": "field1", "error": "required"}]
	Errors []ErrorField `json:"errors"`

	// The TraceID for tracking the error
	//
	// Example: "u8357577235577jdrf9083227"
	TraceID string `json:"traceID"`
}

func SimpleErrorResponse(err error, code string) ErrorResponse {
	return ErrorResponse{
		Error:   err.Error(),
		Code:    code,
		Errors:  nil,
		TraceID: "",
	}
}

func FullErrorResponse(err error, code string, errs []ErrorField) ErrorResponse {
	return ErrorResponse{
		Error:   err.Error(),
		Code:    code,
		Errors:  errs,
		TraceID: "",
	}
}
