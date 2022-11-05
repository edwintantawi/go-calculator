package expconv

import (
	"strings"

	"github.com/edwintantawi/go-calculator/stack"
)

var operatorLevel = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

func ItoP(stmt string) string {
	results := stack.NewStrings()
	operators := stack.NewStrings()

	tokens := strings.Fields(stmt)

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		currentLevel, isOperator := operatorLevel[token]
		if !isOperator {
			results.Push(token)
			continue
		}

		if len(operators) == 0 {
			operators.Push(token)
			continue
		}

		for i := len(operators) - 1; i >= 0; i-- {
			if operatorLevel[operators[i]] < currentLevel || currentLevel == operatorLevel["^"] {
				break
			}

			top := operators.Pop()
			results.Push(top)
		}

		operators.Push(token)
	}

	for len(operators) != 0 {
		topOperator := operators.Pop()
		results.Push(topOperator)
	}

	return strings.Join(results, " ")
}
