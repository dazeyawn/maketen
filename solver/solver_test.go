package solver

import "testing"

func TestSolve(t *testing.T) {
	ok, expr, err := Solve([]int{21, 1, 1, 100}, 21)

	if err != nil {
		t.Fatal("expected true")
	}

	t.Log(ok, expr)
}
