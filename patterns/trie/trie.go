// Package trie contains trie/prefix-tree patterns.
//
// Trie visual after inserting "go" and "goal":
//
//	root
//	 └── g
//	     └── o (word)
//	         └── a
//	             └── l (word)
//
// Shared prefixes are stored once.
package trie

// Trie supports word insertion and prefix lookup.
//
// Common uses:
//   - autocomplete
//   - dictionary search
//   - word search pruning
type Trie struct {
	children map[rune]*Trie
	isWord   bool
}

// New creates an empty trie node.
func New() *Trie {
	return &Trie{children: make(map[rune]*Trie)}
}

// Insert walks or creates one node per character.
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = New()
		}
		node = node.children[ch]
	}
	node.isWord = true
}

// Search requires the full word to exist and end on isWord=true.
func (t *Trie) Search(word string) bool {
	node := t.find(word)
	return node != nil && node.isWord
}

// StartsWith only requires the prefix path to exist.
func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}

// find walks the trie and returns nil if any character is missing.
func (t *Trie) find(s string) *Trie {
	node := t
	for _, ch := range s {
		node = node.children[ch]
		if node == nil {
			return nil
		}
	}
	return node
}
