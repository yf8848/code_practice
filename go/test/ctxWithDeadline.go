package main

import (
	"context"
	"fmt"
	"time"
)

func TestDeadline() {
	d := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
