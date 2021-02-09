package main

import "testing"

func TestGreeting(t *testing.T) {
	expectedValue := "<b>Code.education Rocks!</b>"

	res := greeting("Code.education Rocks!")
	if res != expectedValue {
		t.Errorf("Resultado função greeting: %s; Resultado esperado: %s", res, expectedValue)
	}
}
