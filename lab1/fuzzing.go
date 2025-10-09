package lab1

import (
	"fmt"
	"math/rand"
	"strings"
)

var Alphabet = []rune{'a', 'b'}

var T = []RewritingRule{
	{left: "aabbbaa", right: "abbaba"},
	{left: "abab", right: "aaa"},
	{left: "baaab", right: "abba"},
	{left: "bbbb", right: "ba"},
}

var T_1 = []RewritingRule{
	{left: "bba", right: "bab"},
	{left: "aaaa", right: "aaa"},
	{left: "aaab", right: "aaa"},
	{left: "abaa", right: "aaa"},
	{left: "abab", right: "aaa"},
	{left: "baaa", right: "aaa"},
	{left: "baba", right: "baab"},
	{left: "bbbb", right: "ba"},
	{left: "baaba", right: "aaa"},
	{left: "baabb", right: "aaa"},
	{left: "babbb", right: "baa"},
}

var Graph []*Vertex
var visited = make(map[string]bool)

const maxWordLen = 100
const maxSteps = 100

type RewritingRule struct {
	left, right string
}

type Vertex struct {
	word   string
	parent *Vertex
	way    []string
}

type Queue struct {
	data                        []*Vertex
	capacity, count, head, tail int
}

func InitQueue() *Queue {
	var q Queue
	q.data = make([]*Vertex, 0)
	q.capacity = 0
	q.count = 0
	q.head = 0
	q.tail = 0
	return &q
}

func (p *Queue) Enqueue(x *Vertex) {
	p.data = append(p.data, x)
	p.tail++
	p.count++
	p.capacity++
}

func (p *Queue) Dequeue() *Vertex {
	x := p.data[p.head]
	p.head++
	p.count--
	return x
}

func CountAllEntry(str, s string) []int {
	var indexList []int
	for i := 0; i < len(str)-len(s)+1; i++ {
		relIndex := strings.HasPrefix(str[i:], s)
		if !relIndex {
			continue
		} else {
			indexList = append(indexList, i)
		}
	}
	return indexList
}

func ReplaceFromIndex(word, left, right string, index int) string {
	return word[:index] + strings.Replace(word[index:], left, right, 1)
}

func GenerateWords() string {
	var word strings.Builder
	lenght := rand.Intn(maxWordLen)
	for range lenght {
		word.WriteRune(Alphabet[rand.Intn(2)])
	}
	return word.String()
}

func GenerateChain(word string, srs []RewritingRule) ([]string, string) {
	var chain []string
	steps := rand.Intn(maxSteps)
	currentWord := word
	chain = append(chain, currentWord)

	for range steps {
		orderRules := rand.Perm(len(srs))
		applied := false

		for _, k := range orderRules {
			indexCount := CountAllEntry(currentWord, srs[k].left)
			if len(indexCount) != 0 {
				changes := indexCount[rand.Intn(len(indexCount))]
				currentWord = ReplaceFromIndex(currentWord, srs[k].left, srs[k].right, changes)
				chain = append(chain, currentWord)
				applied = true
				break
			}
		}
		if !applied {
			break
		}
	}

	return chain, currentWord
}

func WordToWord(word, word_1 string, srs_1 []RewritingRule) {
	if word == word_1 {
		fmt.Printf("«%s» and «%s» are equal.\n", word, word_1)
		return
	}
	wordStart, wordFinish := word_1, word
	if word > word_1 {
		wordStart, wordFinish = word, word_1
	}
	chain, _ := BuildTreeBFS(wordStart, wordFinish, srs_1)
	if len(chain) != 0 {
		fmt.Printf("The transition «%s» -> «%s». Chain:%s\n", wordStart, wordFinish, chain)
	} else {
		fmt.Printf("There is no direct chain of transformations from «%s» to «%s»", wordStart, wordFinish)
		chain_1, vertex := BuildTreeBFS2(wordStart, srs_1)
		fmt.Printf("«%s» -> «%s» <- «%s»\n", wordStart, vertex.word, wordFinish)
		fmt.Print(chain)
		fmt.Print(chain_1)
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

func BuildTreeBFS(wordStart, wordFinish string, srs_1 []RewritingRule) ([]string, string) {
	q := InitQueue()
	rootTree := &Vertex{
		word:   wordStart,
		parent: nil,
		way:    nil,
	}

	Graph = append(Graph, rootTree)
	visited[wordStart] = true

	q.Enqueue(rootTree)
	for q.count > 0 {
		currentVertex := q.Dequeue()
		if currentVertex.word == wordFinish {
			return currentVertex.way, currentVertex.word
		}

		newWords := AllVariants(currentVertex.word, srs_1)
		//fmt.Print(currentVertex.word)
		//fmt.Print(newWords)
		//fmt.Print("\n")

		for _, newWord := range newWords {
			if visited[newWord] || len(newWord) < len(wordFinish) {
				continue
			}

			newWay := append(currentVertex.way, currentVertex.word)
			w := &Vertex{
				word:   newWord,
				parent: currentVertex,
				way:    newWay,
			}
			visited[newWord] = true

			Graph = append(Graph, w)
			q.Enqueue(w)
		}
	}
	return nil, ""
}

func BuildTreeBFS2(wordStart string, srs_1 []RewritingRule) ([]string, *Vertex) {
	q := InitQueue()
	rootTree := &Vertex{
		word:   wordStart,
		parent: nil,
		way:    nil,
	}

	q.Enqueue(rootTree)
	for q.count > 0 {
		currentVertex := q.Dequeue()

		newWords := AllVariants(currentVertex.word, srs_1)
		//fmt.Print(currentVertex.word)
		//fmt.Print(newWords)
		//fmt.Print("\n")

		for _, newWord := range newWords {
			if visited[newWord] {
				vertex := FindVertex(newWord)
				return currentVertex.way, vertex
			}

			newWay := append(currentVertex.way, currentVertex.word)
			w := &Vertex{
				word:   newWord,
				parent: currentVertex,
				way:    newWay,
			}
			visited[newWord] = true

			Graph = append(Graph, w)
			q.Enqueue(w)
		}
	}
	return nil, nil
}

func AllVariants(wordStart string, srs_1 []RewritingRule) []string {
	var newStrings []string
	for _, rule := range srs_1 {
		ruleEntry := CountAllEntry(wordStart, rule.left)
		//fmt.Printf("%d %s\n", ruleEntry, rule.left)
		if len(ruleEntry) > 0 {
			for j := 0; j < len(ruleEntry); j++ {
				newStrings = append(newStrings, ReplaceFromIndex(wordStart, rule.left, rule.right, ruleEntry[j]))
			}
		}
	}
	return newStrings
}

func CheckEqualance(word, word_1 string, srs_1 []RewritingRule) {
	if word == word_1 {
		chain_0, NF_0 := GenerateChain(word, srs_1)
		fmt.Printf("«%s» and «%s» are equal.\n", word, word_1)
		fmt.Printf("The normal form of this word: %s\n", NF_0)
		fmt.Printf("The chain of transformations: %s\n", chain_0)
		return
	}
	_, NF := GenerateChain(word, srs_1)
	_, NF_1 := GenerateChain(word_1, srs_1)

	if NF == NF_1 {
		fmt.Printf("Testing for words «%s» and «%s». Fuzzing is successful.\n", word, word_1)
		fmt.Printf("The normal form of this words: %s\n", NF)
	} else {
		fmt.Printf("The SRS T and SRS T_1 are not equivalent.")
		fmt.Printf("The normal form of word «%s»: %s\n", word, NF)
		fmt.Printf("The normal form of word «%s»: %s\n", word_1, NF_1)
	}

	return
}

func main() {
	s := GenerateWords()
	fmt.Print(s + "\n")
	chain, w := GenerateChain(s, T)
	fmt.Print(chain)
	fmt.Print("\n" + w + "\n")
	WordToWord(w, s, T_1)
}
