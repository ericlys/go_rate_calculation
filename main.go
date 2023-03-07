//go mod init packageprojectname //== npm init -y
//go mod tidy //= npm i
//go run main.go

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	fmt.Println(uuid.New().String())
}
