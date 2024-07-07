package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	
	
	if url := os.Getenv("URL"); url != "" {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
		response, err := c.Do(req)
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		fmt.Printf("Body : %s", string(body))
	} else {
		response, err := c.Get("http://localhost:8089/datetime")
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		fmt.Printf("Body : %s", body)
	}
}

//https://www.practical-go-lessons.com/chap-35-build-an-http-client
