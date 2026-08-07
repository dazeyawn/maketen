# MakeTen API

MakeTen API is a simple REST API written in Go that solves arithmetic expression puzzles.

Given 2 to 6 integers and a target value, it searches for an expression using `+`, `-`, `*`, and `/` that evaluates to the target.

This project was created as a learning project for Go, focusing on:

* Building HTTP APIs with the standard library
* JSON encoding and decoding
* Query parameter handling
* Recursive search algorithms
* Writing tests

---

## Features

* REST API using Go's `net/http`
* Accepts 2 to 6 integers
* Supports `+`, `-`, `*`, `/`
* Returns one valid expression if a solution exists
* Includes unit tests for the solver

---

## Requirements

* Go 1.24 or later

---

## Running

```bash
go run .
```

The server starts on:

```
http://localhost:8080
```

---

## API

### Solve

```
GET /solve
```

### Query Parameters

| Name   | Type    | Description              |
| ------ | ------- | ------------------------ |
| nums   | string  | Comma-separated integers |
| target | integer | Target value             |

Example:

```
GET /solve?nums=1,2,3,4&target=10
```

---

## Successful Response

```json
{
  "solvable": true,
  "expression": "(1+2+3+4)"
}
```

If no solution exists:

```json
{
  "solvable": false
}
```

---

## Error Response

```json
{
  "error": "numbers must contain between 2 and 6 values, got 1"
}
```

---

## Running Tests

```
go test ./...
```

---

## Project Structure

```
.
├── main.go
├── solver/
│   ├── solver.go
│   └── solver_test.go
├── go.mod
└── go.sum
```

---

## Future Improvements

* Docker support
* Live reload with Air
* Additional endpoints
* Performance improvements
* Benchmark tests

---

## License

MIT
