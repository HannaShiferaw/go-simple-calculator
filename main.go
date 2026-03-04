package main

import "fmt"

func main() {
    fmt.Println("Advanced Go Calculator")

    for {
        var num1 float64
        var num2 float64
        var op string

        fmt.Print("Enter first number: ")
        fmt.Scanln(&num1)

        fmt.Print("Enter second number: ")
        fmt.Scanln(&num2)

        fmt.Print("Enter operation (+ - * /): ")
        fmt.Scanln(&op)

        

        fmt.Println("Result:", result)

        var cont string
        fmt.Print("Do you want to calculate again? (y/n): ")
        fmt.Scanln(&cont)
        if cont != "y" {
            fmt.Println("Goodbye!")
            break
        }
    }
}
