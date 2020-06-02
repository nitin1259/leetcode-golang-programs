package easy

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GenerateFibonacciSerires function to print the fabonacci series
func GenerateFibonacciSerires() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	checkError(err, "error while reading from standard io")

	input = strings.TrimSpace(input)
	inputNum, err := strconv.Atoi(input)
	checkError(err, "Error while string to int conversion")

	pirntFabonacciSeries(inputNum)

}

func pirntFabonacciSeries(n int) {

	first, second := 0, 1

	fmt.Println(first)
	fmt.Println(second)
	for i := 0; i < n-2; i++ {

		third := first + second
		fmt.Println(third)
		first, second = second, third
	}
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
