package main

import (
	"fmt"
	"math/rand"
	"strings"

	transforms "lab1/transforms"
)

var inv_preservation = make(map[string]string)

func AStay(word, word_1, rule string) {
	if inv_preservation["The 'a' in rule "+rule] != "" {
		return
	}
	word_1_a := len(transforms.CountAllEntry(word_1, "a"))
	if word_1_a == 0 {
		inv_preservation["The 'a' in rule "+rule] = word + " -> " + word_1
	}
}

func ReductionBBB(word, word_1, rule string) {
	if inv_preservation["Reduction of BBB in rule "+rule] != "" {
		return
	}
	if strings.Contains(word, "bbb") {
		word_b3 := len(transforms.CountAllEntry(word, "bbb"))
		word_1_b3 := len(transforms.CountAllEntry(word_1, "bbb"))

		if word_b3-word_1_b3 < 0 {
			inv_preservation["Reduction of BBB in rule "+rule] = word + " -> " + word_1
		}

	}
}

func CheckInvariants(word, word_1, rule string) {
	AStay(word, word_1, rule)
	ReductionBBB(word, word_1, rule)
}

func GenerateMetamorphicChain(minSteps, maxSteps int, word string, srs []transforms.RewritingRule) {
	steps := rand.Intn(maxSteps-minSteps+1) + minSteps
	currentWord := word

	for range steps {
		orderRules := rand.Perm(len(srs))
		applied := false

		for _, k := range orderRules {
			rule := srs[k]
			indexCount := transforms.CountAllEntry(currentWord, rule.Left)
			if len(indexCount) != 0 {
				changes := indexCount[rand.Intn(len(indexCount))]
				leftWord := currentWord
				currentWord = transforms.ReplaceFromIndex(currentWord, rule.Left, rule.Right, changes)
				rightWord := currentWord

				CheckInvariants(leftWord, rightWord, transforms.RewritingRuleToString(rule))
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
	str := transforms.GenerateWords(transforms.MinWordLen, transforms.MaxWordLen)
	GenerateMetamorphicChain(100, 100, str, transforms.T_1)
	fmt.Println("Metamorphic test")

	for inv_rule, example := range inv_preservation {
		fmt.Println("FAIL! " + inv_rule)
		fmt.Printf("Example: %s\n\n", example)
		return
	}

	if len(inv_preservation) == 0 {
		fmt.Println("All invariants are preserved.")
	}
}
