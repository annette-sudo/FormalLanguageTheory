package main

import (
	"fmt"
	"math/rand"
	"strings"

	"lab1/src"
)

var inv_preservation = make(map[string]string)

func ParityA(word, word_1, rule string) {
	if inv_preservation["Parity of A changes in rule "+rule] != "" {
		return
	}
	word_a := len(src.CountAllEntry(word, "a"))
	word_1_a := len(src.CountAllEntry(word_1, "a"))
	if (word_a % 2) != (word_1_a % 2) {
		inv_preservation["Parity of A changes in rule "+rule] = word + " -> " + word_1
	}
}

func ReductionBBB(word, word_1, rule string) {
	if inv_preservation["Reduction of BBB and increase A in rule "+rule] != "" {
		return
	}
	if strings.Contains(word, "bbb") {
		word_a := len(src.CountAllEntry(word, "a"))
		word_1_a := len(src.CountAllEntry(word_1, "a"))
		word_b3 := len(src.CountAllEntry(word, "bbb"))
		word_1_b3 := len(src.CountAllEntry(word_1, "bbb"))
		if (word_b3-word_1_b3 >= 0) && (word_1_a-word_a >= 1) {
			return
		} else {
			inv_preservation["Reduction of BBB and increase A in rule "+rule] = word + " -> " + word_1
		}
	}
}

func AB(word string) int {
	pair := 0
	for i := range word {
		if word[i] == 'a' {
			pair += len(src.CountAllEntry(word[i:], "b"))
		}
	}
	return pair
}

func PairAToTheLeftB(word, word_1, rule string) {
	if inv_preservation["Pairs of A to left B in rule "+rule] != "" {
		return
	}
	word_AB := AB(word)
	word_1_AB := AB(word_1)
	if word_AB < word_1_AB {
		inv_preservation["Pairs of A to left B in rule "+rule] = word + " -> " + word_1
	}
}

func Equation(word, word_1, rule string) {
	if inv_preservation["Equation in rule "+rule] != "" {
		return
	}
	word_aa := len(src.CountAllEntry(word, "aa"))
	word_bb := len(src.CountAllEntry(word, "bb"))
	word_ab := len(src.CountAllEntry(word, "ab"))
	word_ba := len(src.CountAllEntry(word, "ba"))

	word_1_aa := len(src.CountAllEntry(word_1, "aa"))
	word_1_bb := len(src.CountAllEntry(word_1, "bb"))
	word_1_ab := len(src.CountAllEntry(word_1, "ab"))
	word_1_ba := len(src.CountAllEntry(word_1, "ba"))

	eq := word_aa + 2*word_bb - 2*word_ab + 6*word_ba
	eq_1 := word_1_aa + 2*word_1_bb - 2*word_1_ab + 6*word_1_ba

	if eq != eq_1 {
		inv_preservation["Equation in rule "+rule] = word + " -> " + word_1
	}
}

func CheckInvariants(word, word_1, rule string) {
	ParityA(word, word_1, rule)
	ReductionBBB(word, word_1, rule)
	PairAToTheLeftB(word, word_1, rule)
	Equation(word, word_1, rule)
}

func GenerateMetamorphicChain(minSteps, maxSteps int, word string, srs []src.RewritingRule) {
	steps := rand.Intn(maxSteps-minSteps+1) + minSteps
	currentWord := word

	for range steps {
		orderRules := rand.Perm(len(srs))
		applied := false

		for _, k := range orderRules {
			rule := srs[k]
			indexCount := src.CountAllEntry(currentWord, rule.Left)
			if len(indexCount) != 0 {
				changes := indexCount[rand.Intn(len(indexCount))]
				leftWord := currentWord
				currentWord = src.ReplaceFromIndex(currentWord, rule.Left, rule.Right, changes)
				rightWord := currentWord

				CheckInvariants(leftWord, rightWord, src.RewritingRuleToString(rule))
				applied = true
				break
			}
		}
		if !applied {
			break
		}
	}
}

func main() {
	str := src.GenerateWords(src.MinWordLen, src.MaxWordLen)
	GenerateMetamorphicChain(100, 100, str, src.T_1)

	fmt.Println("Metamorphic test")
	for inv_rule, example := range inv_preservation {
		fmt.Println("FAIL!" + inv_rule)
		fmt.Printf("Example: %s\n\n", example)
	}
}
