package responseModel

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
