// Package trie contains trie/prefix-tree patterns.
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

func New() *Trie {
	return &Trie{children: make(map[rune]*Trie)}
}

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

func (t *Trie) Search(word string) bool {
	node := t.find(word)
	return node != nil && node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}

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
