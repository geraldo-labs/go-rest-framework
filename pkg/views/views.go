package views

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/geraldo-labs/go-rest-framework/pkg/serializers"
)

// ViewModel ...
type ViewModel struct {
	ViewModeler
	Serializer serializers.SerializerModel
}

// Post ...
func (v *ViewModel) Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := v.Serializer.SetData(r); err != nil {
		v.Serializer.SendResponse(w, err)
		return
	}
	if err := v.Serializer.Create(); err != nil {
		v.Serializer.SendResponse(w, err)
		return
	}
	v.Serializer.SendResponse(w, v.Serializer.Model)
}

// Delete ...
func (v *ViewModel) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Get ...
func (v *ViewModel) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Head ...
func (v *ViewModel) Head(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Put ...
func (v *ViewModel) Put(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Patch ...
func (v *ViewModel) Patch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// Options ...
func (v *ViewModel) Options(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
