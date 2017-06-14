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

func TestStackPushIncrementsTheSizeOfTheStack(t *testing.T) {
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

func TestStackPushIncreasesTheCapacityByDefaultCapacityIfStorageIsInsufficientForMoreElements(t *testing.T) {
	capacity := 1
	stack := NewStack(capacity)

	data1 := 100
	stack.Push(data1)
	assert.Equal(t, 1, stack.Capacity(), "Expected capacity to be correct")
	assert.Equal(t, 1, stack.Size(), "Expected size to be correct")

	data2 := 200
	stack.Push(data2)
	assert.Equal(t, 2, stack.Size(), "Expected size to be correct")
	assert.Equal(t, DefaultCapacity+capacity, stack.Capacity(), "Expected capacity to be updated")
}

func TestStackPopReturnsTheLatestElementFromTheStack(t *testing.T) {
	capacity := 10
	stack := NewStack(capacity)

	data1 := 100
	data2 := 200
	stack.Push(data1)
	stack.Push(data2)

	value := stack.Pop().(int)

	assert.Equal(t, 1, stack.Size(), "Expected size to be correct")
	assert.Equal(t, data2, value, "Expected latest value to be retreived")
}

func TestStackPopReturnsNilIfStackIsEmpty(t *testing.T) {
	capacity := 10
	stack := NewStack(capacity)

	data1 := 100
	stack.Push(data1)

	value1 := stack.Pop().(int)
	value2 := stack.Pop()

	assert.Equal(t, data1, value1, "Expected latest value to be retreived")
	assert.Nil(t, value2, "Expected value to be nil")
}

func TestIsEmptyReturnsTrueIfTheStackIsEmpty(t *testing.T) {
	capacity := 10
	stack := NewStack(capacity)

	assert.True(t, stack.IsEmpty(), "Expected stack to be empty")
	data1 := 100
	stack.Push(data1)
	stack.Pop()

	assert.True(t, stack.IsEmpty(), "Expected stack to be empty")
}
