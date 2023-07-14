# Functionality

[![Go Report Card](https://goreportcard.com/badge/github.com/vmanikes/Functionality)](https://goreportcard.com/report/github.com/vmanikes/Functionality)
[![Go Reference](https://pkg.go.dev/badge/github.com/vmanikes/Functionality.svg)](https://pkg.go.dev/github.com/vmanikes/Functionality)

## Introduction

Well generics in Go are here and I thought it will be really cool if we could apply the functional programming aspects 
like `map`, `filter` and `reduce` to the slices in our Go apps. 

Functionality lets you stop writing those old, boring `for` loops to do operations on slices

## Stop talking and tell me how it works
Let me show you how :wink:

Lets say you want to filter all even numbers from 1 to 10, square the result and find the sum

The old way to do that is
```go
func Example() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	
	even := make([]int, 0, 0)
	for _, num := range numbers {
		if num % 2 == 0 {
			even = append(even, num)
		}
	}
	
	for idx, num := range even {
		even[idx] = num * num
	}
	
	sum := 0

	for _, num := range even {
		sum += num
	}
}
```

This is how it will look when you use Functionality 
```go
func Functionality() {
	numbers := NewIterator[int]() // Create a new iterator
	
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i) // Seed some data
	}
	
	sum := numbers.Filter(func(num int) bool {
		return num % 2 == 0
	}).Map(func(i int) int {
		return i * i
	}).Reduce(func(accumulator, element int) int {
		return accumulator + element
	}, 0)
	
	fmt.Println("SUM IS ", sum)
}
```

Clean ain't it :D. Well the possibilities are endless and please get creative

## Looks cool, do you use reflection?
Nope, we use generics, and you would need to specify the type when you define the iterator

## What is next?
Well next step is to create some set related operations like, `union`, `intersection` and some other methods

## Looks cool, how can I get involved
Simple, fork the repo, create a change and submit a PR. If you find a bug, please use the issues tab to submit bugs or feature requests
