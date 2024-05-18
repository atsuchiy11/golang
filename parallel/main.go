package parallel

func RunThingsConcurrently() {
	// Create a channel to communicate between goroutines
	c := make(chan int)

	// Start a goroutine that will send a value to the channel
	go func() {
		c <- 42
	}()

	// Read the value from the channel
	value := <-c
	println(value)
}
