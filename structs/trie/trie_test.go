package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieSearch(t *testing.T) {
	testTrie := NewTrie()

	testTrie.Insert("aragorn")
	testTrie.Insert("argon")
	testTrie.Insert("oreo")
	testTrie.Insert("orbit")
	testTrie.Insert("orc")
	testTrie.Insert("orca")

	assert.True(t, testTrie.Search("aragorn"))
	assert.True(t, testTrie.Search("orc"))
	assert.True(t, testTrie.Search("orca"))
	assert.False(t, testTrie.Search("orb"))
}
