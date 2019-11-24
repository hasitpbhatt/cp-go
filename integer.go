package cpgo

// NewInteger creates Integer from int64
func NewInteger(v int64) Integer {
	return Integer{value: v}
}

// GetInteger receives the interface
// and returns int64 (basic int64)
//
// Supports int64, int, Integer as valid input
func GetInteger(v interface{}) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int64:
		return val
	case int32:
		return int64(val)
	case Integer:
		return val.value
	}
	panic("invalid data type")
}

// Integer struct is equivalent to integer
type Integer struct {
	value int64
}

// Less implements the interface required for using the object
// with standard data structures like queue, pair and other similar
// C++ STL like classes
func (i Integer) Less(v interface{}) bool {
	if val, ok := v.(Integer); ok {
		return val.value < i.value
	}
	panic("invalid data type")
}
