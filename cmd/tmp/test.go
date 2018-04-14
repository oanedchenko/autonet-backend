package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	for i:= 1; i < 30; i++ {
		fmt.Printf("%d   %s \r\n", i, uuid.New().String())
	}
}


