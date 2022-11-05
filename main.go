package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/edwintantawi/go-calculator/calculator"
)

func main() {
	fmt.Print(">>> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	stmt := scanner.Text()
	result, err := calculator.Calculate(stmt)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Result:", result)
}
