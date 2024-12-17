package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Ball struct {
	Player string
}

func ping(ctx context.Context, wg *sync.WaitGroup, ball chan Ball) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("game called off")
			return
		case _, ok := <-ball:
			if !ok {
				return
			}
			fmt.Println("ping has the ball")
			time.Sleep(time.Second * 1)
			select {
			case ball <- Ball{"ping"}:
			case <-ctx.Done():
				fmt.Println("game called off")
				return
			}
		}
	}
}

func pong(ctx context.Context, wg *sync.WaitGroup, ball chan Ball) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("game called off")
			return
		case _, ok := <-ball:
			if !ok {
				return
			}
			fmt.Println("pong has the ball")
			time.Sleep(time.Second * 1)
			select {
			case ball <- Ball{"pong"}:
			case <-ctx.Done():
				fmt.Println("game called off")
				return
			}
		}
	}
}
func main() {
	var wg sync.WaitGroup
	d := time.Now().Add(time.Second * 7)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	wg.Add(2)
	ball := make(chan Ball, 1)
	ball <- Ball{"ping"} // Start with "ping" holding the ball
	go ping(ctx, &wg, ball)
	go pong(ctx, &wg, ball)
	wg.Wait()
	close(ball)
	fmt.Println("Game ended and the winner is ", <-ball)
}
