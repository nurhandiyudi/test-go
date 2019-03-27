package main

import (
    "fmt"
    "math"
    )

func main() {
    var X int
    var Y int
    var B int
    
    
    fmt.Printf("Enter X: ")
    fmt.Scanf("%d", &X)
    fmt.Printf("Enter Y: ")
    fmt.Scanf("%d", &Y)
    
    var sum int = sum(X, Y)
	fmt.Printf("Output %d\n", sum)
	
	fmt.Printf("Enter X: ")
    fmt.Scanf("%d", &X)
    fmt.Printf("Enter Y: ")
    fmt.Scanf("%d", &Y)
    
	var multiply int = multiply(X, Y)
	fmt.Printf("Output %d\n", multiply)
	
	fmt.Printf("Enter Jumlah Digit Bilangan Prima: ")
    fmt.Scanf("%d", &B)
    
	j := 1
    for i := 0; j <= B; i++{
       if primecheck(i) {
            fmt.Printf("%v ", i)
            j++
        } 
    }
    
    fmt.Printf("Enter Jumlah Digit Bilangan Prima: ")
    fmt.Scanf("%d", &B)

    for n := 0; n < 4; n++ {
        fmt.Printf("%v ", fibonacci(n))
    }
    
}

func sum(X, Y int) int {
	return X + Y
}

func multiply(X, Y int) int {
	return X * Y
}

func primecheck(digit int) bool {
    for i := 2; i <= int(math.Floor(float64(digit) / 2)); i++ {
        if digit%i == 0 {
            return false
        }
    }
    return digit > 1
}

func fibonacci(n int) int {
    a := 0
    b := 1
    for i := 0; i < n; i++ {
        temp := a
        a = b
        b = temp + a
    }
    return a
}
