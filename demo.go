package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getRand(rander *rand.Rand, max int) int {
	return rander.Intn(max)
}

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	rander := rand.New(seed)
	counts := make(map[int]int)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT)

countLoop:
	for {
		select {
		case _ = <-sigs:
			break countLoop
		default:
			result := getRand(rander, 100)

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
