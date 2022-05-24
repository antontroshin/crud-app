package main

import (
	"context"
	"fmt"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

var client dapr.Client

func main() {
	fmt.Println("starting error generator app")
	go invokeBinding(context.TODO())
	go invokeMethodErrors(context.TODO())
	go invokeMethod(context.TODO())

	// leaving this one in the main goroutine to keep the process alive :)
	publishEvent(context.TODO())
}

func invokeBinding(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("generating invoke binding errors")

			_, err := getDaprClient().InvokeBinding(ctx, &dapr.InvokeBindingRequest{
				Name:      "foo",
				Operation: "baz",
				Data:      []byte("ooo"),
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func invokeMethodErrors(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("generating invoke method errors")

			_, err := getDaprClient().InvokeMethod(ctx, "not-exists", "blah", "get")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func publishEvent(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("generating publish event errors")

			err := getDaprClient().PublishEvent(ctx, "foo", "baz", []byte("data"))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func invokeMethod(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("invoking crud-app")

			data, err := getDaprClient().InvokeMethod(ctx, "crud-app", "/api/v1/todos", "GET")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("invoke crud-app succeeded")
			fmt.Println(string(data))
		}
	}
}

func getDaprClient() dapr.Client {
	if client == nil {
		c, err := dapr.NewClient()
		if err != nil {
			panic(err)
		}
		client = c
	}
	return client
}
