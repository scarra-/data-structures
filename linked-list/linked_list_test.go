package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSingleElement(t *testing.T) {
	llist := LinkedList{}
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 5})
	llist.Prepend(&Node{data: 6})
	llist.Prepend(&Node{data: 4})

	llist.DeleteSingleOccurrenceByValue(5)

	fmt.Println(llist.Values())

	assert.Equal(t, []int{4, 6, 4, 4}, llist.Values())
}

func TestRemoveAllOccurrencesElement(t *testing.T) {
	llist := LinkedList{}
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 5})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 6})
	llist.Prepend(&Node{data: 4})

	llist.DeleteAllOccurrencesByValue(4)

	fmt.Println(llist.Values())

	assert.Equal(t, []int{6, 5}, llist.Values())
}

func TestRemoveAllOccurrencesWithNonExistingElement(t *testing.T) {
	llist := LinkedList{}
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 5})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 6})
	llist.Prepend(&Node{data: 4})

	llist.DeleteAllOccurrencesByValue(100)

	fmt.Println(llist.Values())

	assert.Equal(t, []int{4, 6, 4, 5, 4, 4}, llist.Values())
}

func TestRemoveSingleWithNonExistingElement(t *testing.T) {
	llist := LinkedList{}
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 5})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 6})
	llist.Prepend(&Node{data: 4})

	llist.DeleteSingleOccurrenceByValue(100)

	fmt.Println(llist.Values())

	assert.Equal(t, []int{4, 6, 4, 5, 4, 4}, llist.Values())
}

func TestHasLoop(t *testing.T) {
	llist := LinkedList{}

	cycleNode := &Node{data: 5}

	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(cycleNode)
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 6})
	llist.Prepend(cycleNode)

	assert.True(t, llist.HasLoop())
}

func TestHasNotLoop(t *testing.T) {
	llist := LinkedList{}

	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 4})
	llist.Prepend(&Node{data: 6})

	assert.False(t, llist.HasLoop())
}
