package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		letterIndex := w[i] - 'a'

		if currentNode.children[letterIndex] == nil {
			currentNode.children[letterIndex] = &Node{}
		}

		currentNode = currentNode.children[letterIndex]
	}

	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		letterIndex := w[i] - 'a'

		if currentNode.children[letterIndex] == nil {
			return false
		}

		currentNode = currentNode.children[letterIndex]
	}

	return currentNode.isEnd
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func main() {
	testTrie := NewTrie()

	testTrie.Insert("aragorn")
	ok := testTrie.Search("aragorn")

	fmt.Println(ok)
}
