package shared

func ResetShared() {
	SharedX = 10
	ChannelA = make(chan string)
	ChannelB = make(chan string)
}
