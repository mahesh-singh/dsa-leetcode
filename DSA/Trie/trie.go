package main

import "fmt"

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

func (this *Trie) Search(word string) bool {
	for _, c := range word {
		if this.Children[c-'a'] == nil {
			return false
		}
		this = this.Children[c-'a']
	}
	return this.isTerminal
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

func main() {

	//Your Trie object will be instantiated and called as such:
	obj := Constructor()
	obj.Insert("apple")

	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("banana"))
	fmt.Println(obj.StartsWith("app"))

}
