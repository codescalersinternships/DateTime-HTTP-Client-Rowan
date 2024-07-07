package client

import (
	"fmt"
	"io"
	"net/http"
	// "os"
	"time"
	// "errors"
)

// DateTimeClient is the structure used as an interface to the client request handling
// has the client object and its own GetDateTime method
type DateTimeClient struct {
	client http.Client
}

// ErrWrongParameterPassed expresses the error of passing more than 1 parameter
var ErrWrongParameterPassed error

// NewDateTimeClient is a function used to instaniate a new DateTimeClient struct
func NewDateTimeClient() DateTimeClient {
	return DateTimeClient{}
}

// GetDateTime is a function used by DateTimeClient struct
// It can have 0 parameters or a single parameter to specify URL of server
// It returns []byte which is the response of request
func (dateTimeClient *DateTimeClient) GetDateTime(urlParam ...string) ([]byte, error) {
	dateTimeClient.client = http.Client{Timeout: time.Duration(1) * time.Second}
	var url string
	if len(urlParam) == 0 {
		url = "http://localhost:8089/datetime"
	} else if len(urlParam) > 1 {
		ErrWrongParameterPassed = fmt.Errorf("Too many arguments passed as %v, Only 1 element is required", urlParam)
		return nil, ErrWrongParameterPassed
	} else {
		url = urlParam[0]
	}
	response, err := dateTimeClient.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// func main() {
// client := http.Client{Timeout: time.Duration(1) * time.Second}
// url := os.Getenv("URL")
// if url == "" {
// 	url = "http://localhost:8089/datetime"
// }
// response, err := client.Get(url)
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 		return
// 	}
// 	defer response.Body.Close()
// 	body, _ := io.ReadAll(response.Body)
// 	fmt.Printf("Body : %s", body)
// client := NewDateTimeClient()
// body, _ := client.GetDateTime()
// fmt.Printf("%s\n", body)

// client.GetDateTime()
// // fmt.Println()
// _, err:= client.GetDateTime(os.Getenv("URL"), "rowann")
// fmt.Println(err)
// // fmt.Println()
// client.GetDateTime("http://localhost:8080/datetime")
// fmt.Println()

// }
