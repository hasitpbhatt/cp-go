package cpgo

// NewString creates String object from string
func NewString(v string) String {
	return String{value: v}
}

// GetString receives the interface
// and returns string
//
// Supports string and String as valid input
func GetString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case String:
		return val.value
	}
	panic("invalid data type")
}

// String struct is equivalent to string
type String struct {
	value string
}

// Less implements the interface required for using the object
// with standard data structures like queue, pair and other similar
// C++ STL like classes
func (i String) Less(v interface{}) bool {
	if val, ok := v.(String); ok {
		return val.value < i.value
	}
	panic("invalid data type")
}
