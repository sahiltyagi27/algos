# Trie

Trie stores words by shared prefixes.

Visual:

```text
root
 └── g
     └── o (word)
         └── a
             └── l (word)
```

Good for:

```text
autocomplete
prefix search
dictionary lookup
word search pruning
```

Code:

- [trie.go](trie.go)
