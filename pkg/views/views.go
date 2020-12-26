package views

import (
	"net/http"

	"github.com/geraldo-labs/go-rest-framework/pkg/serializers"
)

// ViewModel ...
type ViewModel struct {
	ViewModeler
	Serializer serializers.SerializerModel
}

// Post ...
func (v *ViewModel) Post(w http.ResponseWriter, r *http.Request) {
	v.Serializer.
}

// Delete ...
func (v *ViewModel) Delete(w http.ResponseWriter, r *http.Request) {

}

// Get ...
func (v *ViewModel) Get(w http.ResponseWriter, r *http.Request) {

}

// Head ...
func (v *ViewModel) Head(w http.ResponseWriter, r *http.Request) {

}

// Put ...
func (v *ViewModel) Put(w http.ResponseWriter, r *http.Request) {

}

// Patch ...
func (v *ViewModel) Patch(w http.ResponseWriter, r *http.Request) {

}

// Options ...
func (v *ViewModel) Options(w http.ResponseWriter, r *http.Request) {

}


