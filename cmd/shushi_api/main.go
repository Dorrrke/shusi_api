package main

import (
	"fmt"

	"github.com/Dorrrke/shusi_api/internal/server"
)

func main() {
	api := server.NewShushiAPI(":8080")
	if err := api.Start(); err != nil {
		fmt.Println(err)
	}
}
