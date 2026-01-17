package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

var Alphabet = []rune{'a', 'b'}

func GenerateWords(minWordLen, maxWordLen int) string {
	var word strings.Builder
	lenght := rand.Intn(maxWordLen-minWordLen+1) + minWordLen
	for range lenght {
		word.WriteRune(Alphabet[rand.Intn(2)])
	}
	return word.String()
}

func generateS() string {
	if rand.Intn(2) == 0 {
		return "b" + generateT() + "aa" + generateT()
	}
	return "ab"
}

func generateT() string {
	r := rand.Intn(3)
	switch r {
	case 0:
		return "a" + generateS() + generateT()
	case 1:
		return "b" + generateT()
	default:
		return "a"
	}
}

type NPDA_State struct {
	order       int
	transitions []*NPDA_Transition
	final       bool
}

type NPDA_Transition struct {
	EndState       *NPDA_State
	symbol         byte
	typeAction     string
	symbolMagazine string
	symbolPut      string
}

type NPDA struct {
	States   []*NPDA_State
	Magazine string
}

func createNPDA() *NPDA {
	q0 := &NPDA_State{order: 0, final: false}
	q1 := &NPDA_State{order: 1, final: false}
	q2 := &NPDA_State{order: 2, final: true}
	q3 := &NPDA_State{order: 3, final: false}
	q4 := &NPDA_State{order: 4, final: false}
	q5 := &NPDA_State{order: 5, final: false}
	q6 := &NPDA_State{order: 6, final: false}
	q7 := &NPDA_State{order: 7, final: false}
	q8 := &NPDA_State{order: 8, final: true}

	q0.transitions = []*NPDA_Transition{
		{EndState: q1, symbol: 'a', typeAction: "NONE"},
		{EndState: q3, symbol: 'b', typeAction: "PUSH", symbolMagazine: "ABA"},
	}

	q1.transitions = []*NPDA_Transition{
		{EndState: q2, symbol: 'b', typeAction: "NONE"},
	}

	q3.transitions = []*NPDA_Transition{
		{EndState: q3, symbol: 'b', typeAction: "NONE"},
		{EndState: q4, symbol: 'a', typeAction: "NONE"},
		{EndState: q6, symbol: 'a', typeAction: "POP", symbolMagazine: "A", symbolPut: ""},
	}

	q4.transitions = []*NPDA_Transition{
		{EndState: q5, symbol: 'a', typeAction: "NONE"},
		{EndState: q3, symbol: 'b', typeAction: "PUSH", symbolMagazine: "ABA"},
	}

	q5.transitions = []*NPDA_Transition{
		{EndState: q3, symbol: 'b', typeAction: "NONE"},
	}

	q6.transitions = []*NPDA_Transition{
		{EndState: q3, symbol: 'e', typeAction: "POP", symbolMagazine: "A", symbolPut: "A"},
		{EndState: q7, symbol: 'a', typeAction: "NONE"},
		{EndState: q8, symbol: 'e', typeAction: "LAST"},
	}

	q7.transitions = []*NPDA_Transition{
		{EndState: q3, symbol: 'a', typeAction: "POP", symbolMagazine: "B", symbolPut: ""},
	}

	NPDA := &NPDA{
		States:   []*NPDA_State{q0, q1, q2, q3, q4, q5, q6},
		Magazine: "",
	}

	return NPDA
}

type NPDA_IntermediateState struct {
	currentWord string
	state       *NPDA_State
	magazine    string
}

func checkNPDA(word string, npda *NPDA) bool {
	accepted := false
	q0 := &NPDA_IntermediateState{
		currentWord: word,
		state:       npda.States[0],
		magazine:    npda.Magazine,
	}

	var stack []*NPDA_IntermediateState
	stack = append(stack, q0)

	for len(stack) != 0 {
		interState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		q := interState.state

		if q.final && len(interState.currentWord) == 0 {
			accepted = true
			break
		}
		var currentSymbol byte
		var newCurrentWord string
		if len(interState.currentWord) != 0 {
			currentSymbol = interState.currentWord[0]
			newCurrentWord = interState.currentWord[1:]
		}
		currentMagazine := interState.magazine
		for i := 0; i < len(q.transitions); i++ {
			trans := q.transitions[i]
			// обработка e-перехода
			if trans.symbol == 'e' && trans.typeAction == "POP" && strings.HasPrefix(interState.magazine, trans.symbolMagazine) {
				newInterState := &NPDA_IntermediateState{
					currentWord: interState.currentWord,
					state:       q.transitions[i].EndState,
					magazine:    trans.symbolPut + strings.TrimPrefix(currentMagazine, trans.symbolMagazine),
				}
				stack = append(stack, newInterState)
			} else if trans.symbol == 'e' && trans.typeAction == "LAST" && interState.magazine == "" {
				newInterState := &NPDA_IntermediateState{
					currentWord: interState.currentWord,
					state:       q.transitions[i].EndState,
					magazine:    "",
				}
				stack = append(stack, newInterState)
			}

			if currentSymbol == trans.symbol {
				if trans.typeAction == "NONE" {
					newInterState := &NPDA_IntermediateState{
						currentWord: newCurrentWord,
						state:       trans.EndState,
						magazine:    interState.magazine,
					}
					stack = append(stack, newInterState)
				} else if trans.typeAction == "PUSH" {
					newInterState := &NPDA_IntermediateState{
						currentWord: newCurrentWord,
						state:       trans.EndState,
						magazine:    trans.symbolMagazine + interState.magazine,
					}
					stack = append(stack, newInterState)
				} else if trans.typeAction == "POP" {
					if strings.HasPrefix(currentMagazine, trans.symbolMagazine) {
						newInterState := &NPDA_IntermediateState{
							currentWord: newCurrentWord,
							state:       trans.EndState,
							magazine:    trans.symbolPut + strings.TrimPrefix(currentMagazine, trans.symbolMagazine),
						}
						stack = append(stack, newInterState)
					}
				}
			}
		}
	}

	return accepted
}

type Grammar struct {
	StartSymbols map[string]bool
	LHS          []string
	Rules        map[string]([][]string)
}

func ParseGrammarFile(filename string) (*Grammar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grammar := &Grammar{
		StartSymbols: make(map[string]bool),
		LHS:          make([]string, 0),
		Rules:        make(map[string]([][]string)),
	}

	scanner := bufio.NewScanner(file)

	startSymbolRegex := regexp.MustCompile(`^<[^>]+>$`)
	ruleRegex := regexp.MustCompile(`^(.+?)\s*->\s*(.+)$`)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if startSymbolRegex.MatchString(line) {
			grammar.StartSymbols[line] = true
			continue
		}

		if matches := ruleRegex.FindStringSubmatch(line); matches != nil {

			left := strings.TrimSpace(matches[1])
			right := strings.TrimSpace(matches[2])
			if len(right) == 1 {
				grammar.Rules[left] = append(grammar.Rules[left], []string{right})
			} else {
				rulesBC := regexp.MustCompile(`^\s*(<[^>]+>)\s+(<[^>]+>)\s*$`)
				if matchesRulesAB := rulesBC.FindStringSubmatch(right); matchesRulesAB != nil {
					B := strings.TrimSpace(matchesRulesAB[1])
					C := strings.TrimSpace(matchesRulesAB[2])
					grammar.Rules[left] = append(grammar.Rules[left], []string{B, C})
				}

			}
		}
	}

	return grammar, scanner.Err()
}

func CYK(word string, grammar Grammar) bool {
	n := len(word)
	if n == 0 {
		return false
	}

	table := make([][]map[string]bool, n)
	for i := range table {
		table[i] = make([]map[string]bool, n+1)
		for j := range table[i] {
			table[i][j] = make(map[string]bool)
		}
	}

	for i := 0; i < n; i++ {
		char := string(word[i])
		for LHS, rules := range grammar.Rules {
			for _, rule := range rules {
				if len(rule) == 1 && rule[0] == char {
					table[i][1][LHS] = true
				}
			}
		}
	}

	for m := 2; m <= n; m++ {
		for i := 0; i <= n-m; i++ {
			for LHS, RHSs := range grammar.Rules {
				for _, rule := range RHSs {
					if len(rule) == 2 {
						B := rule[0]
						C := rule[1]

						for k := 1; k < m; k++ {
							if table[i][k][B] && table[i+k][m-k][C] {
								table[i][m][LHS] = true
								break
							}
						}
					}
				}
			}
		}
	}

	for S := range grammar.StartSymbols {
		if table[0][n][S] {
			return true
		}
	}
	return false
}

func main() {
	grammar, err := ParseGrammarFile("grammar/CF.txt")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	grammar_ll1, err := ParseGrammarFile("grammar/intersection_grammar_LL1.txt")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	grammar_lr0, err := ParseGrammarFile("grammar/intersection_grammar_LR0.txt")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	NPDA := createNPDA()

	fmt.Println("Проверка для слов из языка")
	for i := 0; i < 500; i++ {
		word := generateS()
		if len(word) < 50 {
			cf := CYK(word, *grammar)
			ll1 := CYK(word, *grammar_ll1)
			lr0 := CYK(word, *grammar_lr0)
			pda := checkNPDA(word, NPDA)
			if !(ll1 == lr0 && lr0 == pda && pda == cf) {
				fmt.Println(word)
				fmt.Println("cf:", cf)
				fmt.Println("dpa:", pda)
				fmt.Println("lr0:", lr0)
				fmt.Println("ll1:", ll1)
			}
		}
	}

	fmt.Println("Проверка для всех слов")
	for i := 0; i < 500; i++ {
		word := GenerateWords(3, 50)
		cf := CYK(word, *grammar)
		ll1 := CYK(word, *grammar_ll1)
		lr0 := CYK(word, *grammar_lr0)
		pda := checkNPDA(word, NPDA)
		if !(ll1 == lr0 && lr0 == pda && pda == cf) {
			fmt.Println(word)
			fmt.Println("cf:", cf)
			fmt.Println("dpa:", pda)
			fmt.Println("lr0:", lr0)
			fmt.Println("ll1:", ll1)

		}
	}

}
