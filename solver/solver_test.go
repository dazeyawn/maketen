package solver

import "testing"

func TestSolve(t *testing.T) {
	ok, expr, err := Solve([]int{3, 4, 5, 6, 7}, 21)

	if err != nil {
		t.Fatal("expected true")
	}

	t.Log(ok, expr)
}
