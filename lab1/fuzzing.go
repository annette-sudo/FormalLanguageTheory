package main

import (
	"fmt"
	"slices"

	"lab1/src"
)

var T = []src.RewritingRule{
	{Left: "aabbbaa", Right: "abbaba"},
	{Left: "abab", Right: "aaa"},
	{Left: "baaab", Right: "abba"},
	{Left: "bbbb", Right: "ba"},
}

var T_1 = []src.RewritingRule{
	{Left: "bba", Right: "bab"},
	{Left: "aaaa", Right: "aaa"},
	{Left: "aaab", Right: "aaa"},
	{Left: "abaa", Right: "aaa"},
	{Left: "abab", Right: "aaa"},
	{Left: "baaa", Right: "aaa"},
	{Left: "baba", Right: "baab"},
	{Left: "bbbb", Right: "ba"},
	{Left: "baaba", Right: "aaa"},
	{Left: "baabb", Right: "aaa"},
	{Left: "babbb", Right: "baa"},
}

var Graph []*Vertex
var visited = make(map[string]bool)

const minWordLen = 10
const minSteps = 10

const maxWordLen = 30
const maxSteps = 30

type Vertex struct {
	word   string
	parent *Vertex
}

func WordToWord(word, word_1 string, srs_1 []src.RewritingRule) {
	if word == word_1 {
		fmt.Printf("«%s» and «%s» are equal.\n", word, word_1)
		return
	}

	wordStart, wordFinish := src.ShortLex(word, word_1)

	vertexFinalWord := BuildTreeBFS(wordStart, wordFinish, srs_1)

	if vertexFinalWord != nil {
		chain := BuildChain(vertexFinalWord)
		fmt.Printf("\nThe transition «%s» -> «%s»\n", wordStart, wordFinish)
		fmt.Printf("Rewriting chain performed according to rule T1:\n")
		for _, elChain := range chain {
			fmt.Println(elChain)
		}
		_, NF := src.GenerateChain(100, 100, wordFinish, srs_1)
		fmt.Printf("\nNormal form: %s\n", NF)
	} else {
		fmt.Printf("There is no direct chain of transformations from «%s» to «%s»\n", wordStart, wordFinish)
		commonVertex, lastVertex := BuildTreeBFS2(wordFinish, srs_1)

		chain := BuildChain(commonVertex)
		chain_1 := BuildChain(lastVertex)
		if commonVertex != nil {
			chain_1 = append(chain_1, commonVertex.word)
			fmt.Printf("«%s» -> «%s» <- «%s»\n", wordStart, commonVertex.word, wordFinish)
			fmt.Printf("\nThe rewrating chain «%s» -> «%s»:\n", wordStart, commonVertex.word)

			for _, elChain := range chain {
				fmt.Println(elChain)
			}
			fmt.Printf("\nThe rewrating chain «%s» -> «%s»:\n", wordFinish, commonVertex.word)

			for _, elChain := range chain_1 {
				fmt.Println(elChain)
			}

			_, NF := src.GenerateChain(100, 100, commonVertex.word, srs_1)
			fmt.Printf("\nNormal form: %s\n", NF)
		} else {
			fmt.Println("Let's check equivalence through reduction to normal form")
			_, NF := src.GenerateChain(100, 100, wordStart, srs_1)
			_, NF1 := src.GenerateChain(100, 100, wordFinish, srs_1)
			if NF == NF1 {
				fmt.Println("SRSs are equivalent.")
				fmt.Printf("The normal form of words «%s» and «%s» is «%s»\n", wordStart, wordFinish, NF)
			} else {
				fmt.Println("FAIL! SRSs are not equivalent.")
				fmt.Printf("Example:\n «%s» -> «%s»\n«%s» -> «%s»\n", wordStart, NF, wordFinish, NF1)
			}
		}

	}
}

func FindVertex(word string) *Vertex {
	for _, vertex := range Graph {
		if vertex.word == word {
			return vertex
		}
	}
	return nil
}

func BuildChain(v *Vertex) []string {
	var chain []string
	if v == nil {
		return nil
	}
	for v.parent != nil {
		chain = append(chain, v.word)
		v = v.parent
	}
	chain = append(chain, v.word)
	slices.Reverse(chain)
	return chain
}

func BuildTreeBFS(wordStart, wordFinish string, srs_1 []src.RewritingRule) *Vertex {
	var queue []*Vertex
	rootTree := &Vertex{
		word:   wordStart,
		parent: nil,
	}

	Graph = append(Graph, rootTree)
	queue = append(queue, rootTree)
	visited[wordStart] = true

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		if currentVertex.word == wordFinish {
			return currentVertex
		}

		newWords := src.AllVariantsToRewrite(currentVertex.word, srs_1)

		for _, newWord := range newWords {
			if visited[newWord] || (len(wordFinish) > len(newWord)) {
				continue
			}

			w := &Vertex{
				word:   newWord,
				parent: currentVertex,
			}
			visited[newWord] = true

			Graph = append(Graph, w)
			queue = append(queue, w)
		}
	}
	return nil
}

func BuildTreeBFS2(wordStart string, srs_1 []src.RewritingRule) (*Vertex, *Vertex) {
	var queue []*Vertex

	rootTree := &Vertex{
		word:   wordStart,
		parent: nil,
	}

	queue = append(queue, rootTree)
	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		newWords := src.AllVariantsToRewrite(currentVertex.word, srs_1)

		for _, newWord := range newWords {
			if visited[newWord] {
				commonVertex := FindVertex(newWord)
				return commonVertex, currentVertex
			}

			w := &Vertex{
				word:   newWord,
				parent: currentVertex,
			}
			visited[newWord] = true

			queue = append(queue, w)
		}
	}
	return nil, nil
}

func main() {
	str := src.GenerateWords(minWordLen, maxWordLen)
	fmt.Printf("Generated word: %s\n", str)
	chain, w := src.GenerateChain(minSteps, maxSteps, str, T)
	fmt.Printf("Rewriting chain performed according to rule T:\n")
	for _, newStr := range chain {
		fmt.Println(newStr)
	}
	fmt.Printf("Final word: %s\n\n", w)
	WordToWord(str, w, T_1)
}
