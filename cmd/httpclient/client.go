package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	response, err := c.Get("http://localhost:8080/datetime")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Printf("Body : %s", body)
}
