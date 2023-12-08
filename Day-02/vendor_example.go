// go get github.com/google/uuid
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Generated UUID:", id)
}
