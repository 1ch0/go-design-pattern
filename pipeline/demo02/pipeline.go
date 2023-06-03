package main

import "fmt"

func main() {
	// Set up the pipeline.
	//for n := range Add(Multiply(Generator(10))) { // 3
	//	fmt.Println(n)
	//}
	fmt.Println(<-Sum(Square(Generator(5))))
}

func Generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 0; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func Multiply(in <-chan int) <-chan int { // 1
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func Add(in <-chan int) <-chan int { // 2
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i + i
		}
		close(out)
	}()
	return out
}

func Square(in <-chan int) <-chan int { // 1
	out := make(chan int, 100)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func Sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var Sum int
		for i := range in {
			Sum += i
		}
		out <- Sum
		close(out)
	}()
	return out
}
