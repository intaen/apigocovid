package domain

type (
	// Response example swagger
	SuccessResponse struct {
		StatusCode string      `json:"code" swaggertype:"string" example:"000"`
		Message    string      `json:"message" example:"[message]"`
		Result     interface{} `json:"result" swaggertype:"object"`
	}
	BadRequestResponse struct {
		StatusCode string      `json:"code" swaggertype:"string" example:"004"`
		Message    string      `json:"message" example:"[message]"`
		Result     interface{} `json:"result" swaggertype:"string" example:"null"`
	}
	ServerErrorResponse struct {
		StatusCode string      `json:"code" swaggertype:"string" example:"005"`
		Message    string      `json:"message" example:"[message]"`
		Result     interface{} `json:"result" swaggertype:"string" example:"null"`
	}
	UnauthorizedErrorResponse struct {
		StatusCode string      `json:"code" swaggertype:"string" example:"041"`
		Message    string      `json:"message" example:"[message]"`
		Result     interface{} `json:"result" swaggertype:"string" example:"null"`
	}

	// Input example swagger
	ExamplePredictName struct {
		Name string `json:"name" example:"Admin"`
	}
)
