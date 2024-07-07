package main

import (
	"fmt"
	// "os"
	"github.com/codescalersinternships/DateTime-HTTP-Client-Rowan/pkg"
)

func main() {
	client := pkg.NewDateTimeClient("http://localhost:8080")
	data, _ := client.GetDateTime()
	fmt.Printf("%s", data)
}
