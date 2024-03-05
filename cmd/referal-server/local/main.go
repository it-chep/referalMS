package main

import (
	"context"
	"log"
	"referalMS/internal"
)

func main() {
	ctx := context.Background()
	log.Fatal(internal.NewApp(ctx).Run())
}
