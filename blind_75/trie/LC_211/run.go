package main

import "fmt"

func main() {

	wordSearchDictionary := Constructor()

	wordSearchDictionary.AddWord("day")
	wordSearchDictionary.AddWord("bay")
	wordSearchDictionary.AddWord("may")
	fmt.Println(wordSearchDictionary.Search("say")) // return false
	fmt.Println(wordSearchDictionary.Search("day")) // return true
	fmt.Println(wordSearchDictionary.Search(".ay")) // return true
	fmt.Println(wordSearchDictionary.Search("b..")) // return true
}