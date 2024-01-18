package respones

type CustomResponse struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}
