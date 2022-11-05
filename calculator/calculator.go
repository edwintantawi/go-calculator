package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/edwintantawi/go-calculator/expconv"
	"github.com/edwintantawi/go-calculator/stack"
)

func Calculate(stmt string) (float64, error) {
	stmt = strings.TrimSpace(stmt)
	if stmt == "" {
		return 0, errors.New("statement cannot be empty")
	}

	c := expconv.ItoP(stmt)
	tokens := strings.Fields(c)

	results := stack.NewFloats()
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(results) < 2 {
				return 0, errors.New("missing required operand")
			}

			right := results.Pop()
			left := results.Pop()
			switch token {
			case "+":
				results.Push(left + right)
			case "-":
				results.Push(left - right)
			case "*":
				results.Push(left * right)
			case "/":
				results.Push(left / right)
			case "^":
				results.Push(math.Pow(left, right))
			}
		default:
			fToken, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("invalid operand")
			}
			results.Push(fToken)
		}
	}

	return results[0], nil
}
