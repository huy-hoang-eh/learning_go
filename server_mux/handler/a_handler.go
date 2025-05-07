package handler

import (
	"fmt"
	"net/http"

	"server_mux/shared"
)

func HandleA(w http.ResponseWriter, r *http.Request) {
	fmt.Println("A Start")
	shared.ChannelA <- "A Start"

	shared.SharedX *= 10
	fmt.Println("A: first sharedX = ", shared.SharedX)
	shared.SharedX += 10
	fmt.Println("A: second sharedX = ", shared.SharedX)

	<-shared.ChannelB
	fmt.Println("A End")
}
