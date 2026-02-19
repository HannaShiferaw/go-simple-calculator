package main

import 
 "fmt"

func main() {
 fmt.Println("Simple Go Calculator: Add two numbers")
    
    var num1 int
    var num2 int
    
    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)
    
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)
    
    fmt.Println("Result:", num1+num2)
}