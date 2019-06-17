package uuidarray

import "github.com/google/uuid"

// RemoveDuplicates removes duplicate uuid records
func (x *UUIDArray) RemoveDuplicates() {
	m := make(map[uuid.UUID]struct{})
	newList := make(UUIDArray, 0)
	for _, v := range *x {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			newList = append(newList, v)
		}
	}
	*x = newList
}
