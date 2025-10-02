package main

import "fmt"

/*
Simple Hello World program in Go
*/

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("This is a simple Go program.")
    
    // Simple calculation
    result := 2 + 2
    fmt.Printf("2 + 2 = %d\n", result)
    
    // Simple slice operation
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Printf("Sum of %v = %d\n", numbers, sum)
    
    // Simple function call
    greeting := greet("World")
    fmt.Println(greeting)
}

func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}