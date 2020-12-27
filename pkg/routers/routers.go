package routers

import "github.com/julienschmidt/httprouter"

func NewHttpRouter() *httprouter.Router {
	return httprouter.New()
}
