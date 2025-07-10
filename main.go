package main

import "fmt"

var Configuration Config

func main() {
	Configuration = readConfig()
	c := make(chan int64)

	go awaitHour(c)
	go initDiscord(c)

	fmt.Println("Started routines")
	select {}
}
