package main

import (
	"fmt"
	"github.com/websbytodd/go-starter/internal/config"
)

func main() {
	props := config.ReadProperties()
	fmt.Printf("secret = %s", props.JWTSecret)
}
