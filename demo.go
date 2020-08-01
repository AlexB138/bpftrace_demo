package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getRand(max int) int {
	return rand.Intn(max)
}

func main() {
	counts := make(map[int]int)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

countLoop:
	for {
		select {
		case _ = <-sigs:
			break countLoop
		default:
			result := getRand(100)

			if _, ok := counts[result]; !ok {
				counts[result] = 0
			}

			counts[result]++

			fmt.Println(result)
			time.Sleep(250 * time.Millisecond)
		}
	}

	fmt.Println(counts)
}
