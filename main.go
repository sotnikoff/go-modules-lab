package main

import (
	"fmt"

	"modlab/mathx"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewV7()
	fmt.Println(id, err)
	fmt.Println(uuid.NewString())
	fmt.Println(mathx.Double(21))
}
