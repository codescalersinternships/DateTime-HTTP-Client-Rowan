package pkg

import (
	"io"
	"net/http"
	"time"
)

// DateTimeClient is the structure used as an interface to the client request handling
// has the client object and its own GetDateTime method
type DateTimeClient struct {
	client http.Client
	url    string
}

// // ErrWrongParameterPassed expresses the error of passing more than 1 parameter
// var ErrWrongParameterPassed error

// NewDateTimeClient is a function used to instaniate a new DateTimeClient struct
func NewDateTimeClient(url string) DateTimeClient {
	return DateTimeClient{url: url + "/datetime"}
}

// GetDateTime is a function used by DateTimeClient struct
// It can have 0 parameters or a single parameter to specify URL of server
// It returns []byte which is the response of request
func (dateTimeClient *DateTimeClient) GetDateTime() ([]byte, error) {
	dateTimeClient.client = http.Client{Timeout: time.Duration(1) * time.Second}
	return retry(dateTimeClient.requestInfo, 3)
}

func (dateTimeClient *DateTimeClient) requestInfo() ([]byte, error) {
	response, err := dateTimeClient.client.Get(dateTimeClient.url)
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

func retry(f func() ([]byte, error), retryAttempts int) ([]byte, error) {
	var err error
	var body []byte
	for i := 0; i < retryAttempts; i++ {
		body, err = f()
		if err == nil {
			return body, err
		}
		time.Sleep(1 * time.Second)

	}
	return body, err
}
