package main

import (
	"fmt"
	"sso/internal/config"

	ssov1 "protos/gen/go/sso"
)

func main() {
	req := &ssov1.RegisterRequest{
		Email:    "user@example.com",
		Password: "password123",
	}

	fmt.Printf("Request: %+v\n", req)

	config := config.MustLoad()
	fmt.Printf("Config: %+v\n", config)
}
