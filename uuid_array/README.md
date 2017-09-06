# uuid array

refer to [github.com/lib/pq](https://github.com/lib/pq)

and using uuid package with [github.com/google/uuid](https://github.com/google/uuid)

## using

```go
type TestStruct struct {
	UUIDs uuidarray.UUIDArray `db:"ids"`
}

func (x *TestStruct) GetUUIDs() []uuid.UUID {
	return []uuid.UUID(x.UUIDs)
}
```
