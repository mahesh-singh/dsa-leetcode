package main

import "fmt"

/*
Implement Trie ll
Implement Trie DS with following function additional functions apart from Insert, Search and StartWith
1. CountWordEqualTo
2. CountWordStartWith
3. Erase
*/

type Trie struct {
	Children      [26]*Trie
	CountWord     int
	WordStartWith int
}

func NewTrie() Trie {
	return Trie{}
}

func (curr *Trie) Insert(word string) {

	/*
		for _, char := range word {
			node := curr.Children[char-'a']
			if node == nil {
				node = &Trie{}
			}

			node.WordStartWith = node.WordStartWith + 1
			curr = node
		}
		curr.CountWord = curr.CountWord + 1
	*/

	for _, c := range word {
		if curr.Children[c-'a'] == nil {

			curr.Children[c-'a'] = &Trie{}
		}
		curr.Children[c-'a'].WordStartWith = curr.Children[c-'a'].WordStartWith + 1
		curr = curr.Children[c-'a']
	}

	curr.CountWord = curr.CountWord + 1

}

func (curr *Trie) CountWordEqualTo(word string) int {
	for _, char := range word {
		node := curr.Children[char-'a']
		if node == nil {
			return -1
		}
		curr = node
	}
	return curr.CountWord
}

func (curr *Trie) CountWordStartWith(word string) int {
	for _, char := range word {
		node := curr.Children[char-'a']
		if node == nil {
			return -1
		}
		curr = node
	}
	return curr.WordStartWith
}

func main() {
	trie := NewTrie()

	trie.Insert("apple")
	trie.Insert("apps")
	trie.Insert("app")
	fmt.Println(trie.CountWordEqualTo("apple"))

	fmt.Println(trie.CountWordStartWith("ap"))
}
