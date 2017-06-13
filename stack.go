package almost

// DefaultCapacity is the default value for the stack
const DefaultCapacity = 10

// Value type for the stack elements.
type _value struct {
	value interface{}
}

// Creates a Value obect for the underlying value.
func NewValue(value interface{}) _value {
	return _value{
		value: value,
	}
}

// Stack represents the stack datastructure.
type Stack struct {
	data []_value
	size int
}

// NewStack creates a new instance of the stack with the given capacity.
// Capacity value < 0 will default to the DefaultCapacity
func NewStack(capacity int) *Stack {
	if capacity <= 0 {
		capacity = DefaultCapacity
	}
	return &Stack{
		data: make([]_value, capacity),
		size: 0,
	}
}

//Capacity provides the underlying capacity of the stack. This is the amount of elements that it can hold
func (s *Stack) Capacity() int {
	return len(s.data)
}

// Size gives the length of elemets that has been pushed to the stack
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(v interface{}) {
	s.data[s.size] = NewValue(v)
	s.size++
}

func (s *Stack) Peek() interface{} {
	return s.data[s.size-1].value
}
