package serializers

import (
	"github.com/geraldo-labs/go-rest-framework/pkg/views"
	"net/http"
)

// Serializerer ...
type Serializerer interface {
	GetData() *views.APIError
	SetData(*http.Request) *views.APIError
	Validate() *views.APIError
	SendResponse(w http.ResponseWriter, data interface{})
	Create() *views.APIError
}
