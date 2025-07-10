package main

import (
	"fmt"
	"strconv"
	"time"
)

func awaitHour(c chan int64) {
	untilMin := 59 - time.Now().Minute()
	untilSec := 59 - time.Now().Second()

	strDur := strconv.Itoa(untilMin) + "m" + strconv.Itoa(untilSec+1) + "s"

	dur, err := time.ParseDuration(strDur)
	if err != nil {
		fmt.Printf("Failed to parse time until top of hour: %s\n", err)
		return
	}

	t := time.NewTimer(dur)
	fmt.Printf("Will tick in: %s\n", strDur)
	begin := <-t.C

	tick(c)
	fmt.Printf("Began ticking at: %s\n", begin)
}

func tick(c chan int64) {
	c <- time.Now().Unix()
	fmt.Printf("First tick at: %s\n", time.Now())
	t := time.NewTicker(1 * time.Hour)

	go func() {
		for {
			tick := <-t.C
			c <- tick.Unix()
			fmt.Printf("Ticked at: %s\n", tick)
		}
	}()
}
