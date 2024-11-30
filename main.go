package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	val, err := fetchUser(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
	fmt.Println("Took:", val)
}

func fetchUser(ctx context.Context, userId int) (int, error) {
	val, err := callThirdPartyDependency()
	if err != nil {
		return 0, err
	}
	return val, err
}

func callThirdPartyDependency() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 665, nil
}
