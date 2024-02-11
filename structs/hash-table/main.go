package main

import "fmt"

const ArraySize = 10

type HashTable struct {
	array [ArraySize]*Bucket
}

func NewHashTable() *HashTable {
	table := &HashTable{array: [ArraySize]*Bucket{}}

	for i := 0; i < ArraySize; i++ {
		table.array[i] = &Bucket{}
	}

	return table
}

type Bucket struct {
	head *Node
}

type Node struct {
	value string
	next  *Node
}

func (t *HashTable) Insert(key string) {
	index := hash(key)
	t.array[index].Insert(key)
}

func (t *HashTable) Search(key string) bool {
	index := hash(key)
	return t.array[index].Search(key)
}

func (t *HashTable) Delete(key string) bool {
	index := hash(key)

	return t.array[index].Delete(key)
}

func (b *Bucket) Insert(w string) {
	if b.head == nil {
		b.head = &Node{value: w}
		return
	}

	currentNode := b.head

	for currentNode != nil {
		if currentNode.value == w {
			fmt.Println("SAME")
			return
		}

		currentNode = currentNode.next
	}

	newHead := &Node{value: w, next: b.head}
	b.head = newHead
}

func (t *Bucket) Search(key string) bool {
	currentNode := t.head

	for currentNode != nil {
		if currentNode.value == key {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (t *Bucket) Delete(key string) bool {
	if t.head == nil {
		return false
	}

	deleted := false

	if t.head.value == key {
		t.head = t.head.next
		return true
	}

	currentNode := t.head

	for currentNode.next != nil {
		if currentNode.next.value == key {
			currentNode.next = currentNode.next.next
			deleted = true
			continue
		}

		currentNode = currentNode.next
	}

	return deleted
}

func hash(key string) int {
	keyLen := len(key)
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % keyLen
}

func main() {
	hashTable := NewHashTable()

	hashTable.Insert("ERIC")
	hashTable.Insert("KYLE")
	hashTable.Insert("STAN")
	hashTable.Insert("KENNY")

	fmt.Println(hashTable.Search("ERIC"))
	fmt.Println(hashTable.Delete("ERIC"))
	fmt.Println(hashTable.Search("ERIC"))

	fmt.Println(hashTable.Search("KYLE"))
	fmt.Println(hashTable.Delete("KYLE"))
	fmt.Println(hashTable.Search("KYLE"))
}
