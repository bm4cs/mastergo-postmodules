package _1_getting_started

import (
	"fmt"
	"os"
	"strconv"
)

/*
Write a program that plays the game of FizzBuzz. The rules are simple:

Count from 1 to a given number n.
Print out each number, with the following exceptions:
If the number is divisible by 3, print “Fizz” instead of the number.
If the number is divisible by 5, print “Buzz”.
If the number is divisible by 15, print “FizzBuzz”.
*/

func fizzbuzz(n int) {
	for i:=1; i <= n; i++ {
		switch {
		case i % 15 == 0:
			fmt.Print("FizzBuzz ")
		case i % 5 == 0:
			fmt.Print("Buzz ")
		case i % 3 == 0:
			fmt.Print("Fizz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
}

func FizzBuzzCli() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
