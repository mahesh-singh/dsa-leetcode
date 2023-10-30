package main

type WordDictionary struct {
	Children [26]*WordDictionary
	IsWord   bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	for _, c := range word {
		if this.Children[c-'a'] == nil {
			this.Children[c-'a'] = &WordDictionary{}
		}
		this = this.Children[c-'a']
	}
	this.IsWord = true
}

func (this *WordDictionary) Search(word string) bool {
	for i, c := range word {
		if c == '.' {
			for _, child := range this.Children {
				if child != nil && child.Search(word[i+1:]) {
					return true
				}
				// return false //making mistake as return will happen without running entire loop
			}
			return false

		} else if this.Children[c-'a'] == nil {
			return false
		}
		this = this.Children[c-'a']
	}
	return this.IsWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
