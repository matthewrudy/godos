package models

import (
	"testing"
)

func Test_FindAll(t *testing.T) {
	todos := FindAll()

	if todos[0].ID != "alph" {
		t.Error("aint alpha", todos[0])
	}
}
