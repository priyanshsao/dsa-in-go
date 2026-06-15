package main 

type WdNode struct {
	children [26]*WdNode
	end bool
}

func New() *WdNode {

	return new(WdNode)
}

type WordSearchDictionary struct {
	root *WdNode
}

func Constructor() *WordSearchDictionary {

	return &WordSearchDictionary{root: New()}
}

func (wsd *WordSearchDictionary) AddWord(word string) {

	curr := wsd.root

	for _, r := range word {
		index := r - 'a'

		if curr.children[index] == nil {
			curr.children[index] = New()
		}
		curr = curr.children[index]
	}
	curr.end = true
}

func (wsd *WordSearchDictionary) Search(word string) bool {

	return wsd.dfs(word, 0, wsd.root)
}

func (wsd *WordSearchDictionary) dfs(word string, j int, root *WdNode) bool {

	curr := root

	for i:=j;i<len(word);i++ {

		c := word[i]

		if c == '.' {

			for _, child := range curr.children {

				if child != nil && wsd.dfs(word,i+1,child) {
					return true
				}
			}
			return false

		} else {
			index := c-'a'

			if curr.children[index] == nil {
				return false
			}
			curr = curr.children[index]
		}
	}
	
	return curr.end
}