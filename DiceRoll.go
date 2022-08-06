package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"os"
	"bufio"
	"strings"
)

func roll(die int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(die) + 1
	return result
}

func PrintSli(sli []int) {
	for i, v := range sli {
		fmt.Printf("%d: %d\n", i + 1, v)
	}
}

func SumSli(sli []int) int {
	sum := 0
	for _, v := range sli {
		sum = sum + v
	}
	return sum
}

func NoArgs() {
	fmt.Printf("Number of dice to roll: ")
	var numstr string
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan()
	input1 := strings.Fields(scanner1.Text())
	if len(input1) > 1 {
		fmt.Printf("Please only enter a single integer number for the number of dice to roll.\n")
		return
	}
	numstr = input1[0]
	num, err := strconv.Atoi(numstr)
	if err != nil {
		fmt.Printf("'%s' is not a number. Please try again with your desired number of dice to roll.\n", numstr)
		return
	}
	fmt.Printf("Type of die to roll: ")
	var diestr string
	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Scan()
	input2 := strings.Fields(scanner2.Text())
	if len(input2) > 1 {
		fmt.Printf("Please only enter a single integer number for the type of dice to roll.\n")
		return
	}
	diestr = input2[0]
	die, err := strconv.Atoi(diestr)
	if err != nil {
		fmt.Printf("'%s' is not a number. Please try again with your desired number of sides per die.\n", diestr)
		return
	}
	results := make([]int, num)
	for i := 0; i < num; i++ {
		res := roll(die)
		results[i] = res
	}
	PrintSli(results)
	fmt.Printf("Total: %d\n", SumSli(results))
}

func HasArgs() {
	var numstr string
	numstr = os.Args[1]
	num, err := strconv.Atoi(numstr)
	if err != nil {
		fmt.Printf("'%s' is not a number. Please try again with your desired number of dice to roll.\n", numstr)
		return
	}
	var diestr string
	diestr = os.Args[2]
	die, err := strconv.Atoi(diestr)
	if err != nil {
		fmt.Printf("'%s' is not a number. Please try again with your desired number of sides per die.\n", diestr)
		return
	}
	results := make([]int, num)
	for i := 0; i < num; i++ {
		res := roll(die)
		results[i] = res
	}
	PrintSli(results)
	fmt.Printf("Total: %d\n", SumSli(results))
}

func main() {
	if len(os.Args[1:]) == 0 {
		NoArgs()
	} else if len(os.Args[1:]) == 1 {
		fmt.Printf("Please provide arguments in format 'roll [number of dice] [type of dice]' or no arguments for a menu.\n")
	} else if len(os.Args[1:]) == 2 {
		HasArgs()
	} else {
		fmt.Printf("Please provide arguments in format 'roll [number of dice] [type of dice]' or no arguments for a menu.\n")
	}
	
}