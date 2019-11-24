package cpgo

// NewBool creates Bool from bool
func NewBool(v bool) Bool {
	return Bool{value: v}
}

// GetBool receives the interface
// and returns bool
//
// Supports bool, and Bool as valid input
func GetBool(v interface{}) bool {
	switch val := v.(type) {
	case bool:
		return val
	case Bool:
		return val.value
	}
	panic("invalid data type")
}

// Bool struct is equivalent to integer
type Bool struct {
	value bool
}
