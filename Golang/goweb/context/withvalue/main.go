package main

import (
	"context"
	"fmt"
)

func main() {

	//extract value of a context withValue
	value := SetWithValue("api-key")
	fmt.Println(value)
}

func SetWithValue(key string) (result any) {
	ctx := context.WithValue(context.Background(), "api-key", "super-secret-api-key")
	result = ctx.Value(key)
	return
}
