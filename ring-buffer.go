package main

import "fmt"

type RingBuffer struct {
	inputCh  <-chan int
	outputCh chan int
}

func (r *RingBuffer) Run() {
	for v := range r.inputCh {
		select {
		case r.outputCh <- v:
		default:
			<-r.outputCh
			r.outputCh <- v
		}
	}
	close(r.outputCh)
}

func main() {
	in := make(chan int)
	out := make(chan int, 5)
	rb := &RingBuffer{in, out}
	go rb.Run()

	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in)

	for res := range out {
		fmt.Println(res)
	}
}
