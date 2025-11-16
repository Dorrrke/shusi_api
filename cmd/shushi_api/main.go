package main

import (
	"fmt"

	"github.com/Dorrrke/shusi_api/internal/server"
)

func main() {
	srv := server.NewShushiAPI(":8080")
	if err := srv.Start(); err != nil {
		fmt.Println(err)
	}
}
