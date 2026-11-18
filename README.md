# Go Simple Calculator

A small calculator that runs in the terminal. It adds, subtracts,
multiplies and divides two numbers, and keeps asking until you
choose to stop.

## Run it

You need Go 1.25 or newer.

```sh
go run .
```

Or build it first, then run the binary:

```sh
go build -o go-simple-calculator .
./go-simple-calculator
```

## Example

```
Go Simple Calculator
--------------------
Enter first number: 10
Enter second number: 4
Enter operation (+ - * /): /
Result: 2.5
Do you want to calculate again? (y/n): n
Goodbye!
```

Invalid numbers and operators are rejected and you are asked again.
Division by zero is refused with an error message instead of a
strange result.

## Tests

```sh
go test ./...
go vet ./...
```

## Files

- `main.go` — the calculator and the input prompts. The maths lives
  in `calculate()`, separate from reading input, so it can be tested.
- `main_test.go` — table tests for all four operations, division by
  zero, and unknown operators.
- `go.mod` — the Go module, no outside dependencies.
