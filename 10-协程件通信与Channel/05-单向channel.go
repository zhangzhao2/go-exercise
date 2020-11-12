package main

func send(out chan<- int) {
	out <- 89
	close(out)
}

func main() {
	ch := make(chan int)
	go func() {
		send(ch)
	}()
}



