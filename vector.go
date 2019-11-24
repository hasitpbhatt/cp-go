package cpgo

// Vector is equivalent to an array with functionalities like
// push, pop etc.
type Vector struct {
	vector []interface{}
}

// NewVector returns the vector of interfaces
func NewVector(n int) *Vector {
	return &Vector{
		vector: make([]interface{}, 0, n),
	}
}

// PushBack pushes the item at the end of the vector
func (v *Vector) PushBack(input interface{}) {
	v.vector = append(v.vector, input)
}

// PopBack pops the item at the end of the vector
func (v *Vector) PopBack(input interface{}) {
	v.vector = v.vector[:len(v.vector)-1]
}

// PushFront inserts the item at the start of the vector
func (v *Vector) PushFront(input interface{}) {
	v.vector = append([]interface{}, v.vector...)
}

// PopFront pops the item at the beginning of the vector
func (v *Vector) PopFront(input interface{}) {
	v.vector = v.vector[1:]
}

// Sort sorts the vector based on the value
//
// l: 0-based start index
// r: 1-based end index 
func (v *Vector) Sort(l,r int) error {
	sort.Slice(v.vector[l:r], func(i,j int) bool {
		return v.vector[l+i].(Comparable).Less(v.vector[l+j])
	})
}
