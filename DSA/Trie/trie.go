package main

import (
	"fmt"
)

type Trie struct {
	Children   [26]*Trie
	isTerminal bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {

	for _, c := range word {
		if this.Children[c-'a'] == nil {

			this.Children[c-'a'] = &Trie{}
		}
		this = this.Children[c-'a']
	}
	this.isTerminal = true
}

func (this *Trie) InsertRecursive(word string) {
	insertRecursive(this, word, 0)
}

func insertRecursive(curr *Trie, word string, i int) {
	if i == len(word) {
		curr.isTerminal = true
		return
	}
	if curr.Children[word[i]-'a'] == nil {
		curr.Children[word[i]-'a'] = &Trie{}
	}
	curr = curr.Children[word[i]-'a']

	insertRecursive(curr, word, i+1)
}

func (this *Trie) Search(word string) bool {
	for _, c := range word {
		if this.Children[c-'a'] == nil {
			return false
		}
		this = this.Children[c-'a']
	}
	return this.isTerminal
}

func (this *Trie) SearchRecursive(word string) bool {
	return searchRecursive(this, word, 0)
}

func searchRecursive(curr *Trie, word string, i int) bool {
	if i == len(word) {
		return curr.isTerminal
	}
	curr = curr.Children[word[i]-'a']
	if curr == nil {
		return false
	}

	return searchRecursive(curr, word, i+1)
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.Children[c-'a'] == nil {
			return false
		}
		this = this.Children[c-'a']
	}
	return true
}

func (this *Trie) Delete(word string) {

}

func main() {

	//Your Trie object will be instantiated and called as such:
	obj := Constructor()
	obj.Insert("apple")

	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("banana"))
	fmt.Println(obj.StartsWith("app"))

}
