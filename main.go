package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	val, err := fetchUser(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
	fmt.Println("Took:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUser(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respChan := make(chan Response)

	go func() {
		val, err := callThirdPartyDependency()
		respChan <- Response{
			value: val,
			err:   err,
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data took too long")
		case resp := <-respChan:
			return resp.value, resp.err
		}
	}
}

func callThirdPartyDependency() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 665, nil
}
