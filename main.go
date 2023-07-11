package main

import (
	"context"
	"fmt"
)

type key string

var myKey key = "Hola mundo"

func Background() {
	fmt.Printf("Im brackground!")
}

func DoSomething(ctx context.Context) {
	fmt.Printf("DoSomething: miKey's value is %s\n", ctx.Value(myKey))
	ctxAnother := context.WithValue(ctx, myKey, "Another value")
	DoAnother(ctxAnother)
	fmt.Printf("DoSomething: miKey's value is %s\n", ctx.Value(myKey))
}

func DoAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value(myKey))
}

func main() {

	// ctx := context.TODO()
	ctx := context.Background()
	ctx = context.WithValue(ctx, myKey, "myValue")

	DoSomething(ctx)

}
