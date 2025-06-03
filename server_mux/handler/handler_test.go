package handler_test

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	channelA := make(chan string)

	go func() {
		channelA <- "A"
		channelA <- "B"
		channelA <- "B"
		time.Sleep(13 * time.Second)
		channelA <- "B"

		//close(channelA) -> 0x99
	}()

	for i := 0; i < 4; i++ {
		fmt.Println(<-channelA)
	}
}

func TestGoRoutineState2(t *testing.T) {
	channelA := make(chan string)
	go func() {
		fmt.Println("Value: ", <-channelA)
	}()

	time.Sleep(10 * time.Second)
	// can not detect deadlock, since main thread is killed -> leading to mem leaked
}

func TestGoRoutineState3(t *testing.T) {
	channelA := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
	}()

	<-channelA
	// deadlock detected, but after 10s to detect the deadlock
	// assumption: we could assume that when doing <-channelA, main thread is switching to waiting state
	// after 10s the number of waiting thread = the number of current thread -> deadlock
	// cons: late deadlock detection
}

func TestGoRoutineState4(t *testing.T) {
	// Prove assumption in state3
	channelA := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("10s go routine finishes")
	}()

	go func() {
		time.Sleep(20 * time.Second)
		fmt.Println("20s go routine finishes	")
	}()

	<-channelA
	// need 20s to detect deadlock
}

func TestGoRoutineState5(t *testing.T) {
	// Classic deadlock pattern
	ch1 := make(chan int)
	ch2 := make(chan int)

	ch1 <- <-ch2
}

func TestGoRoutineState(t *testing.T) {
	channelA := make(chan string)
	go func() {
		fmt.Println("Value: ", <-channelA)
	}()

	time.Sleep(10 * time.Second)
	channelA <- "Hello"
	close(channelA)
}
