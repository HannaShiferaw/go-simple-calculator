// Command go-simple-calculator is a small interactive calculator
// that runs in the terminal. It adds, subtracts, multiplies and
// divides two numbers, and keeps asking until you choose to stop.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculate applies op to num1 and num2.
// It returns an error for an unknown operator and for division by zero.
func calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operation %q: use + - * or /", op)
	}
}

// readLine reads one line from the terminal and trims surrounding spaces.
func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

// readNumber asks for a number until the user types a valid one.
func readNumber(reader *bufio.Reader, prompt string) (float64, error) {
	for {
		fmt.Print(prompt)
		text, err := readLine(reader)
		if err != nil {
			return 0, err
		}
		num, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("That is not a valid number. Please try again.")
			continue
		}
		return num, nil
	}
}

// readOperator asks for an operator until the user types + - * or /.
func readOperator(reader *bufio.Reader) (string, error) {
	for {
		fmt.Print("Enter operation (+ - * /): ")
		op, err := readLine(reader)
		if err != nil {
			return "", err
		}
		switch op {
		case "+", "-", "*", "/":
			return op, nil
		default:
			fmt.Println("Unknown operation. Please use + - * or /.")
		}
	}
}

func main() {
	fmt.Println("Go Simple Calculator")
	fmt.Println("--------------------")

	reader := bufio.NewReader(os.Stdin)

	for {
		num1, err := readNumber(reader, "Enter first number: ")
		if err != nil {
			fmt.Println("Goodbye!")
			return
		}

		num2, err := readNumber(reader, "Enter second number: ")
		if err != nil {
			fmt.Println("Goodbye!")
			return
		}

		op, err := readOperator(reader)
		if err != nil {
			fmt.Println("Goodbye!")
			return
		}

		result, err := calculate(num1, num2, op)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}

		fmt.Print("Do you want to calculate again? (y/n): ")
		cont, err := readLine(reader)
		if err != nil || (cont != "y" && cont != "Y") {
			fmt.Println("Goodbye!")
			return
		}
	}
}
