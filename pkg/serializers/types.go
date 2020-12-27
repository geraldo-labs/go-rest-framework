package serializers

type APIError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}
