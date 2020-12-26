package serializers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

// SerializerModel ...
type SerializerModel struct {
	Serializerer
	Model gorm.Model
}

// Data ...
func (s *SerializerModel) Data(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&s.Model)
	if err != nil {
		return err
	}
	return nil
}

// Validate ...
func (s *SerializerModel) Validate() error {
	err := json.NewDecoder(r.Body).Decode(&s.Model)
	if err != nil {
		return err
	}
	return nil
}
