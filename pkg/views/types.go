package views

type APIError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}
