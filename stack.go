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

// Push adds the element to the top of the stack.
func (s *Stack) Push(v interface{}) {
	//if size reaches capacity, increase the capacity by the default limit
	if s.size == s.Capacity() {
		data := make([]_value, s.size+DefaultCapacity)
		copy(data, s.data)
		s.data = data
	}
	s.data[s.size] = NewValue(v)
	s.size++
}

// Peek retreives the latest element without removing it from the stack.
func (s *Stack) Peek() interface{} {
	return s.data[s.size-1].value
}

// Pop retreives the latest element and removes it from the stack.
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		s.size--
		return s.data[s.size].value
	}
	return nil
}

// IsEmpty returns true if stack is empty, false otherwise
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
