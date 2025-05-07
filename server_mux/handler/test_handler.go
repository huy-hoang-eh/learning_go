package handler

import (
	"fmt"
	"net/http"
	"sync"

	"server_mux/shared"
)

func testA(wg *sync.WaitGroup, w http.ResponseWriter, r *http.Request) {
	_, err := http.Get("http://localhost:8080/a")

	if err != nil {
		fmt.Println(err)
	}
	defer wg.Done()
}

func testB(wg *sync.WaitGroup, w http.ResponseWriter, r *http.Request) {
	_, err := http.Get("http://localhost:8080/b")

	if err != nil {
		fmt.Println(err)
	}
	defer wg.Done()
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------------------")
	shared.ResetShared()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go testA(&wg, w, r)
	go testB(&wg, w, r)

	wg.Wait()
	fmt.Println("---------------------------")
}
