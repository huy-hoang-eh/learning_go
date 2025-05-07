package handler

import (
	"fmt"
	"net/http"

	"server_mux/shared"
)

func HandleB(w http.ResponseWriter, r *http.Request) {
	<-shared.ChannelA
	fmt.Println("B Start")

	shared.SharedX += 10
	fmt.Println("B: first sharedX = ", shared.SharedX)
	shared.SharedX *= 10
	fmt.Println("B: second sharedX = ", shared.SharedX)

	fmt.Println("B End")
	shared.ChannelB <- "B End"
}
