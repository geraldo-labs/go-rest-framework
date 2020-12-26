package model

type Modeler interface {
	Validate() error
	Save() error
}
