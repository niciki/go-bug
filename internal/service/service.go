package service

import (
	"context"
	"fmt"
	"time"
)

type FouService struct {
}

func (s FouService) PrintString(ctx context.Context, id string) (err error) {
	go PrintStringGoRoutine(ctx, id)
	return
}

func PrintStringGoRoutine(ctx context.Context, id string) {
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("curr id: %s\n", id)
	}
}

func (s FouService) Get(ctx context.Context, id string) (str string, err error) {
	str = "answer"
	time.Sleep(1 * time.Second)
	return
}
