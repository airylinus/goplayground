package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.Tick(3 * time.Second)
	for {
		<-timer
		fmt.Println("done")
	}
}
