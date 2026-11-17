package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewV7()
	fmt.Println(id, err)
	fmt.Println(uuid.NewString())
}
