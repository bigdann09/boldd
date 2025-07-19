package dtos

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}
