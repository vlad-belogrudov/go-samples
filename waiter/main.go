// Example of expired clock function
package main

import (
	"context"
	"fmt"
	"time"
)

// Clock runs pseudo-graphic until expiration
func Clock(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Printf("\r[  #  ]")
			return
		default:
			for _, c := range `-\|/` {
				fmt.Printf("\r[  %c  ]", c)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	go Clock(ctx)
	fmt.Scanln()
}
