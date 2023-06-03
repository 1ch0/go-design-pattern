package main

func main() {
	for out := range Dry(Wash(Buy(10))) {
		println(out)
	}
}

func Buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
		}
	}()
	return out
}

func Wash(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "Wash-" + c
		}
	}()
	return out
}

func Dry(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "Dry-" + c
		}
	}()
	return out
}
