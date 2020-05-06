package test

import(
	"testing"
)

func TestHolaMundo(t *testing.T) {
	str := "Hola Mundo"
	if str != "Hola Mundo" {
		t.Error("No es posible saludar a los usuarios", nil)
	}
}