package main

import (
	"fmt"

	"modlab/mathx"

	"github.com/google/uuid"
	"github.com/sotnikoff/greeter"
)

func main() {
	id, err := uuid.NewV7()
	fmt.Println(id, err)
	fmt.Println(uuid.NewString())
	fmt.Println(mathx.Double(21))
	fmt.Println(greeter.Hello("Михаил"))
}
