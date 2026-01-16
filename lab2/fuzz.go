package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var Alphabet = []rune{'a', 'b'}

const MaxLen = 100

func generateWord() string {
	var word strings.Builder
	lenght := rand.Intn(MaxLen + 1)
	for range lenght {
		word.WriteRune(Alphabet[rand.Intn(2)])
	}
	return word.String()
}

func checkRegex(word string) bool {
	part1 := `^(a|b)*abb(a|b)$`
	part2 := `^(a|ba)*bb(a*b)*aba(a|b)$`
	re1 := regexp.MustCompile(part1)
	re2 := regexp.MustCompile(part2)
	return re1.MatchString(word) || re2.MatchString(word)
}

func checkDFA(word string) bool {
	det := [][]int{
		{0, 14},
		{5, 1},
		{9, 10},
		{0, 4},
		{0, 1},
		{6, 7},
		{6, 13},
		{8, 2},
		{11, 12},
		{6, 7},
		{5, 1},
		{6, 13},
		{8, 2},
		{5, 2},
		{0, 2},
	}

	Final := map[int]bool{
		9:  true,
		10: true,
		11: true,
		12: true,
	}

	q := 3
	str := []rune(word)

	for i := 0; i < len(word); i++ {
		var transition int
		symbol := str[i]
		if symbol == 'a' {
			transition = 0
		} else {
			transition = 1
		}
		q = det[q][transition]
	}

	if Final[q] {
		return true
	}

	return false

}

type NFA_State struct {
	order       int
	transitions []*NFA_Transition
	final       bool
}

type NFA_Transition struct {
	EndState *NFA_State
	symbol   byte
}

type NFA struct {
	States []*NFA_State
}

type NFA_IntermediateState struct {
	currentWord string
	state       *NFA_State
}

func createNFA() *NFA {
	q0 := &NFA_State{order: 0, final: false}
	q1 := &NFA_State{order: 1, final: false}
	q2 := &NFA_State{order: 2, final: false}
	q3 := &NFA_State{order: 3, final: false}
	q4 := &NFA_State{order: 4, final: false}
	q5 := &NFA_State{order: 5, final: false}
	q6 := &NFA_State{order: 6, final: false}
	q7 := &NFA_State{order: 7, final: false}
	q8 := &NFA_State{order: 8, final: false}
	q9 := &NFA_State{order: 9, final: true}

	q0.transitions = []*NFA_Transition{
		{EndState: q0, symbol: 'a'},
		{EndState: q1, symbol: 'a'},
		{EndState: q0, symbol: 'b'},
		{EndState: q2, symbol: 'b'},
	}

	q1.transitions = []*NFA_Transition{
		{EndState: q4, symbol: 'b'},
	}

	q2.transitions = []*NFA_Transition{
		{EndState: q3, symbol: 'b'},
	}

	q3.transitions = []*NFA_Transition{
		{EndState: q5, symbol: 'a'},
		{EndState: q6, symbol: 'a'},
	}

	q4.transitions = []*NFA_Transition{
		{EndState: q7, symbol: 'b'},
	}

	q5.transitions = []*NFA_Transition{
		{EndState: q5, symbol: 'a'},
		{EndState: q3, symbol: 'b'},
	}

	q6.transitions = []*NFA_Transition{
		{EndState: q8, symbol: 'b'},
	}

	q7.transitions = []*NFA_Transition{
		{EndState: q9, symbol: 'b'},
		{EndState: q9, symbol: 'a'},
	}

	q8.transitions = []*NFA_Transition{
		{EndState: q7, symbol: 'a'},
	}

	nfa := &NFA{
		States: []*NFA_State{q0, q1, q2, q3, q4, q5, q6, q7, q8, q9},
	}

	return nfa
}

func checkNFA(word string, nfa *NFA) bool {
	accepted := false
	q0 := &NFA_IntermediateState{
		currentWord: word,
		state:       nfa.States[0],
	}

	var stack []*NFA_IntermediateState
	stack = append(stack, q0)

	for len(stack) != 0 {
		interState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		q := interState.state

		if q.final && len(interState.currentWord) == 0 {
			accepted = true
			break
		} else if len(interState.currentWord) == 0 {
			continue
		}
		currentSymbol := interState.currentWord[0]
		newCurrentWord := interState.currentWord[1:]
		for i := 0; i < len(q.transitions); i++ {
			if currentSymbol == q.transitions[i].symbol {
				newInterState := &NFA_IntermediateState{
					currentWord: newCurrentWord,
					state:       q.transitions[i].EndState,
				}
				stack = append(stack, newInterState)

			}
		}
	}

	return accepted
}

type AFA_State struct {
	order       int
	transitions []*AFA_Transition
	final       bool
}

type AFA_Transition struct {
	EndStates     []*AFA_State
	symbol        byte
	isConjunctive bool
}

type AFA struct {
	States []*AFA_State
}

type AFA_IntermediateState struct {
	currentWord string
	states      []*AFA_State
}

func createAFA() *AFA {
	q0 := &AFA_State{order: 0, final: false}
	q1 := &AFA_State{order: 1, final: false}
	q2 := &AFA_State{order: 2, final: false}
	q3 := &AFA_State{order: 3, final: false}
	q4 := &AFA_State{order: 4, final: true}
	q5 := &AFA_State{order: 5, final: false}
	q6 := &AFA_State{order: 6, final: false}
	q7 := &AFA_State{order: 7, final: false}
	q8 := &AFA_State{order: 8, final: false}
	q9 := &AFA_State{order: 9, final: false}
	q10 := &AFA_State{order: 10, final: false}

	q0.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q0}, symbol: 'a', isConjunctive: false},
		{EndStates: []*AFA_State{q0}, symbol: 'b', isConjunctive: false},
		{EndStates: []*AFA_State{q1}, symbol: 'a', isConjunctive: false},

		{EndStates: []*AFA_State{q5, q7}, symbol: 'b', isConjunctive: true},
	}

	q1.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q2}, symbol: 'b', isConjunctive: false},
	}

	q2.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q3}, symbol: 'b', isConjunctive: false},
	}

	q3.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q4}, symbol: 'a', isConjunctive: false},
		{EndStates: []*AFA_State{q4}, symbol: 'b', isConjunctive: false},
	}

	q5.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q6}, symbol: 'b', isConjunctive: false},
	}

	q6.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q6}, symbol: 'a', isConjunctive: false},
		{EndStates: []*AFA_State{q6}, symbol: 'b', isConjunctive: false},
		{EndStates: []*AFA_State{q4}, symbol: 'a', isConjunctive: false},
		{EndStates: []*AFA_State{q4}, symbol: 'b', isConjunctive: false},
	}

	q7.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q7}, symbol: 'a', isConjunctive: false},
		{EndStates: []*AFA_State{q7}, symbol: 'b', isConjunctive: false},
		{EndStates: []*AFA_State{q8}, symbol: 'b', isConjunctive: false},
	}

	q8.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q9}, symbol: 'a', isConjunctive: false},
	}

	q9.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q10}, symbol: 'b', isConjunctive: false},
	}

	q10.transitions = []*AFA_Transition{
		{EndStates: []*AFA_State{q3}, symbol: 'a', isConjunctive: false},
	}

	afa := &AFA{
		States: []*AFA_State{q0, q1, q2, q3, q4, q5, q6, q7, q8, q9, q10},
	}

	return afa
}

func checkAFA(qStart *AFA_State, word string) bool {
	accepted := false
	q0 := &AFA_IntermediateState{
		currentWord: word,
		states:      []*AFA_State{qStart},
	}

	var stack []*AFA_IntermediateState
	stack = append(stack, q0)

	for len(stack) != 0 {
		interState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		qs := interState.states
		if len(qs) == 1 {
			q := qs[0]
			if q.final && len(interState.currentWord) == 0 {
				accepted = true
				break
			} else if len(interState.currentWord) == 0 {
				continue
			}
			currentSymbol := interState.currentWord[0]
			newCurrentWord := interState.currentWord[1:]
			for i := 0; i < len(q.transitions); i++ {
				if currentSymbol == q.transitions[i].symbol {
					newInterState := &AFA_IntermediateState{
						currentWord: newCurrentWord,
						states:      q.transitions[i].EndStates,
					}
					stack = append(stack, newInterState)

				}
			}
		} else if len(qs) == 2 {
			q1 := qs[0]
			q2 := qs[1]

			if checkAFA(q1, interState.currentWord) && checkAFA(q2, interState.currentWord) {
				return true
			}
		}
	}

	return accepted
}

func main() {
	NFA := createNFA()
	AFA := createAFA()
	equivalent := true

	for range 100000 {
		s := generateWord()
		check := checkRegex(s)
		if !(check == checkDFA(s) && check == checkNFA(s, NFA) && checkAFA(AFA.States[0], s) == check) {
			equivalent = false
			fmt.Printf("Word: %s/nRegex: %t/nDFA: %t/nNFA: %t/nAFA: %t/n", s, check, checkDFA(s), checkNFA(s, NFA), checkAFA(AFA.States[0], s))
		}

	}
	fmt.Println()
	fmt.Println(equivalent)

}
