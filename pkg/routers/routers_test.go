package routers

import (
	"testing"
)

func TestNewHttpRouter(t *testing.T) {
	router := NewHttpRouter()
	if router == nil {
		t.Errorf("NewHttpRouter must not be nil: %v", router)
	}
}
