package serializers

import (
	"net/http"
)

// Serializerer ...
type Serializerer interface {
	GetData() *APIError
	SetData(*http.Request) *APIError
	Validate() *APIError
	SendResponse(w http.ResponseWriter, data interface{})
	Create() *APIError
}
