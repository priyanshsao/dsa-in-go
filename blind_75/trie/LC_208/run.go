package main

import (
	"fmt"

	"github.com/priyanshsao/dsa-in-go/base/trie"
)

func main() {

	trie := trie.Constructor()

	trie.Insert("apple")

	fmt.Println(trie.Search("apple"))// should return true
	fmt.Println(trie.Search("app")) // should return false

	fmt.Println(trie.StartsWith("app")) // should return true

	trie.Insert("app")
	fmt.Println(trie.Search("app")) // should return true
}