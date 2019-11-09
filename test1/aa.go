package main

import (
	"context"
	"fmt"
	"github.com/golang_samples/http/utils"
	"time"
)

func main() {
	utils.SayHello()
}

func ctx() {
	ctx, cancel := context.WithCancel(context.Background())

	// monitor
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-ctx.Done():
				fmt.Println("return")
				return
			default:
				fmt.Println("monitor woring")

			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
