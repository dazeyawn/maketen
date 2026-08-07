package solver

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type DecimalOp func(a, b decimal.Decimal) decimal.Decimal

var operators = map[string]DecimalOp{
	"+": func(a, b decimal.Decimal) decimal.Decimal {
		return a.Add(b)
	},
	"-": func(a, b decimal.Decimal) decimal.Decimal {
		return a.Sub(b)
	},
	"*": func(a, b decimal.Decimal) decimal.Decimal {
		return a.Mul(b)
	},
	"/": func(a, b decimal.Decimal) decimal.Decimal {
		return a.Div(b)
	},
}

func Solve(numbers []int, target int) (bool, string, error) {
	nLen := len(numbers)

	if nLen < 2 || nLen > 6 {
		return false, "", fmt.Errorf("numbers must contain between 2 and 6 values, got %d", nLen)
	}

	decimals := make([]decimal.Decimal, nLen)
	sNums := make([]string, nLen)
	for i := range nLen {
		decimals[i] = decimal.NewFromInt(int64(numbers[i]))
		sNums[i] = strconv.Itoa(numbers[i])
	}
	t := decimal.NewFromInt(int64(target))
	stack := make([]decimal.Decimal, 0, nLen)
	fStack := make([]string, 0, nLen*2-1)
	used := make([]bool, nLen)
	var dfs func(depth int) (bool, string)
	dfs = func(depth int) (bool, string) {
		if depth == nLen*2-1 {
			if stack[0].Equal(t) {
				return true, fStack[0]
			}
			return false, ""
		}

		for i := range decimals {
			if !used[i] {
				used[i] = true
				stack = append(stack, decimals[i])
				fStack = append(fStack, sNums[i])

				if ok, expr := dfs(depth + 1); ok {
					return true, expr
				}

				used[i] = false
				stack = stack[:len(stack)-1]
				fStack = fStack[:len(fStack)-1]
			}
		}

		if len(stack) >= 2 {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			f1 := fStack[len(fStack)-2]
			f2 := fStack[len(fStack)-1]
			for key, op := range operators {
				if key == "/" && b.IsZero() {
					continue
				}
				if (key == "+" || key == "*") && a.GreaterThan(b) {
					continue
				}
				stack = stack[:len(stack)-2]
				stack = append(stack, op(a, b))
				fStack = fStack[:len(fStack)-2]
				fStack = append(fStack, fmt.Sprintf("(%s%s%s)", f1, key, f2))

				if ok, expr := dfs(depth + 1); ok {
					return true, expr
				}

				stack = append(stack[:len(stack)-1], a, b)
				fStack = append(fStack[:len(fStack)-1], f1, f2)
			}
		}

		return false, ""
	}

	found, expr := dfs(0)
	return found, expr, nil
}
