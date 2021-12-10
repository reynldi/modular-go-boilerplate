package response

import "github.com/google/uuid"

type ResponseTemplate struct {
	TraceID string `json:"traceId"`
	Status  int    `json:"status"`
}

type SuccessResponse struct {
	ResponseTemplate
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func FormatSuccessResponse(status int, message string, result interface{}) SuccessResponse {
	response := SuccessResponse{
		ResponseTemplate: ResponseTemplate{
			TraceID: uuid.New().String(),
			Status:  status,
		},
		Message: message,
		Result:  result,
	}

	return response
}

type ErrorResponse struct {
	ResponseTemplate
	Status int         `json:"status"`
	Errors interface{} `json:"errors"`
}

func FormatErrorResponse(status int, message string, errors interface{}) ErrorResponse {
	response := ErrorResponse{
		ResponseTemplate: ResponseTemplate{
			TraceID: uuid.New().String(),
			Status:  status,
		},
		Status: status,
		Errors: errors,
	}

	return response
}
