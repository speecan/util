package uuidarray

import (
	"testing"

	"github.com/google/uuid"
)

func TestValue(t *testing.T) {
	x := UUIDArray{uuid.New()}
	res, err := x.Value()
	if err != nil {
		t.Fatal(err)
	}
	uuid.New().Value()
	_ = res
}
