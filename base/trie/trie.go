// This file contains code related to tries.
//
// The implementation uses go map.
//
// Defines three methods on trie: Insert, Search, Prefix.
package trie


type TrieNode struct {
	children map[rune]*TrieNode
	endNode  bool
}

type Trie struct {
	Root *TrieNode
}

// New creates a new trie node.
func New() *TrieNode {

	tn := new(TrieNode)
	tn.children = make(map[rune]*TrieNode)
	
	return tn
}

// GetChildren returns the children of a trie node.
func (tn *TrieNode) GetChildren() map[rune]*TrieNode {

	return tn.children
}

// IsEnd returns true if the trie node is the end of series,
// otherwise false.
func (tn *TrieNode) IsEnd() bool {

	return tn.endNode
}

// Constructor creates a new Trie data structure.
func Constructor() *Trie {

	t := new(Trie)
	t.Root = New()

	return t
} 

// Insert inserts the input in Trie data structure.
func (t *Trie) Insert(word string) {
	curr := t.Root

	for _, r := range word {
		if curr.children[r] == nil {
			curr.children[r] = New()
		}
		curr = curr.children[r]
	}
	curr.endNode = true
}

// Search searches the input in Trie.
func (t *Trie) Search(word string) bool {
	curr := t.Root

	for _, r := range word {
		if curr.children[r] == nil {
			return false
		}

		curr = curr.children[r]
	}

	return curr.endNode
}

// StartsWith checks if the input is present in any series.
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.Root

	for _, r := range prefix {
		if curr.children[r] == nil {
			return false
		}

		curr = curr.children[r]
	}
	
	return true
}