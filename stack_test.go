package almost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// NewStack
func TestNewStackCreatesStackWithTheProvidedCapacity(t *testing.T) {
	capacity := 20

	stack := NewStack(capacity)

	assert.Equal(t, capacity, stack.Capacity(), "Expected stack to be created with the provided capacity")
}

func TestNewStackCreatesStackWithDefaultCapacityWhenProvidedWithInvalidCapacity(t *testing.T) {
	capacity := -1

	stack := NewStack(capacity)

	assert.Equal(t, DefaultCapacity, stack.Capacity(), "Expected stack to be created with the provided capacity")
}

func TestNewStackCreatesStackCreateStackWithZeroSize(t *testing.T) {
	capacity := 10

	stack := NewStack(capacity)

	assert.Equal(t, 0, stack.Size(), "Expected stack to be created with the provided capacity")
}

func TestStackPushUpdatesTheSizeOfTheStack(t *testing.T) {
	capacity := 10
	stack := NewStack(capacity)

	data1 := 100
	stack.Push(data1)
	assert.Equal(t, 1, stack.Size(), "Expected stack to be created with the provided capacity")

	data2 := 200
	stack.Push(data2)
	assert.Equal(t, 2, stack.Size(), "Expected stack to be created with the provided capacity")
}

func TestStackPushStoresAnElementInTheTopOfTheList(t *testing.T) {
	capacity := 10
	stack := NewStack(capacity)

	data1 := 100
	stack.Push(data1)
	value := stack.Peek().(int)
	assert.Equal(t, 1, stack.Size(), "Expected stack to be created with the provided capacity")
	assert.Equal(t, data1, value, "Expected stack to be created with the provided capacity")

	data2 := 200
	stack.Push(data2)
	value = stack.Peek().(int)
	assert.Equal(t, 2, stack.Size(), "Expected stack to be created with the provided capacity")
	assert.Equal(t, data2, value, "Expected stack to be created with the provided capacity")
}
