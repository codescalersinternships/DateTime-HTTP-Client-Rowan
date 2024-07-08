package main

import (
	"fmt"
	"github.com/codescalersinternships/DateTime-HTTP-Client-Rowan/pkg"
	"os"
)

func main() {
	client := pkg.NewDateTimeClient(os.Getenv("URL"))
	// client := pkg.NewDateTimeClient("http://localhost:8080")
	data, _ := client.GetDateTime()
	fmt.Printf("%s", data)
	fmt.Println(string(data))
}
