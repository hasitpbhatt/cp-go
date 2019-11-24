package cpgo

// NewDouble creates Double object from float64
func NewDouble(v float64) Double {
	return Double{value: v}
}

// GetDouble receives the interface
// and returns float64
//
// Supports float64 and Double as valid input
func GetDouble(v interface{}) float64 {
	switch val := v.(type) {
	case float64:
		return val
	case Double:
		return val.value
	}
	panic("invalid data type")
}

// Double struct is equivalent to float64
type Double struct {
	value float64
}

// Less implements the interface required for using the object
// with standard data structures like queue, pair and other similar
// C++ STL like classes
func (i Double) Less(v interface{}) bool {
	if val, ok := v.(Double); ok {
		return val.value < i.value
	}
	panic("invalid data type")
}
