package uuidarray

import (
	"testing"

	"github.com/google/uuid"
)

func TestRemoveDuplicate(t *testing.T) {
	u1 := uuid.New()
	u2 := uuid.New()
	value := UUIDArray{u1, u2, uuid.MustParse(u1.String()), uuid.MustParse(u1.String())}
	if len(value) != 4 {
		t.Fatal("failed to prepare uuid array data, length must be 4, acctual:", len(value))
	}
	value.RemoveDuplicates()
	if len(value) != 2 {
		t.Fatal("failed to remove duplicate data from uuid array data, length must be 2, acctual:", len(value))
	}
}

func BenchmarkRemoveDuplicate(b *testing.B) {
	v := make(UUIDArray, 2^20)
	for i := 0; i < len(v); i++ {
		v[i] = uuid.New()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.RemoveDuplicates() // its almost unique array
	}
}
