package main

import (
	"strings"
	"golang.org/x/tour/wc"
)


func WordCount(s string) map[string]int {
	words := strings.Split(s," ")
	ok := false
	m := make(map[string]int)

	for i:=0;i<len(words);i++{
		_,ok = m[words[i]
		if ok{
			m[words[i]]=m[words[i]]+1
		}
		else{
			m[words[i]] = 1
		}


	}
	return m
}

func main() {
	wc.Test(WordCount)
}

