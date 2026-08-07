package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dazeyawn/maketen/solver"
)

type SolveRequest struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type SolveResponse struct {
	Solvable   bool   `json:"solvable"`
	Expression string `json:"expression,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func SolveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()

	nums := query.Get("nums")
	target := query.Get("target")

	if nums == "" {
		http.Error(w, "missing nums", http.StatusBadRequest)
		return
	}

	if target == "" {
		http.Error(w, "missing target", http.StatusBadRequest)
		return
	}

	parts := strings.Split(nums, ",")
	numbers := make([]int, len(parts))
	for i, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "invalid nums", http.StatusBadRequest)
			return
		}
		numbers[i] = n
	}

	t, err := strconv.Atoi(target)
	if err != nil {
		http.Error(w, "invalid target", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	solvable, expression, err := solver.Solve(numbers, t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{
			Error: err.Error(),
		}); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	response := SolveResponse{
		Solvable:   solvable,
		Expression: expression,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
