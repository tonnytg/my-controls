package main

import (
	"context"
	"fmt"
	"log"

	acm "cloud.google.com/go/accesscontextmanager/apiv1"
)

func main() {
	log.Println("Start VPC Service Control")

	acc := acm.NewClient(context.Background())

	acc.GetAccessLevel()
	fmt.Println()
}
