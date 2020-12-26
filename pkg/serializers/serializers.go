package serializers

import (
	"encoding/json"
	"fmt"
	"github.com/geraldo-labs/go-rest-framework/pkg/model"
	"github.com/geraldo-labs/go-rest-framework/pkg/views"
	"log"
	"net/http"
)

// SerializerModel ...
type SerializerModel struct {
	Serializer
	Model model.Modeler
}

// Data Receives data from payload and parse into the model
func (s *SerializerModel) SetData(r *http.Request) *views.APIError {
	err := json.NewDecoder(r.Body).Decode(&s.Model)
	if err != nil {
		return &views.APIError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("%s", fmt.Errorf("%w", err)),
		}
	}
	return nil
}

// Validate Run validation on the model.
func (s *SerializerModel) Validate() *views.APIError {
	if err := s.Model.Validate(); err != nil {
		return &views.APIError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("%s", fmt.Errorf("%w", err)),
		}
	}
	return nil
}

// GetData Get serialized data from the model
func (s *SerializerModel) GetData() *views.APIError {
	if err := s.Model.Validate(); err != nil {
		return &views.APIError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("%s", fmt.Errorf("%w", err)),
		}
	}
	return nil
}

// Create Creates data using registered model
func (s *SerializerModel) Create() *views.APIError {
	return nil
}

type Serializer struct {
}

// SendResponse Send response JSON properly formatted
func (s *Serializer) SendResponse(w http.ResponseWriter, data interface{}) {
	err2 := json.NewEncoder(w).Encode(data)
	if err2 != nil {
		log.Printf("error encoding response: %v", fmt.Errorf("%w", err2))
		w.WriteHeader(500)
	}
}
