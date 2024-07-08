package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_GetDateTime(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, time.Now())
	}))
	defer mockServer.Close()
	client := NewDateTimeClient(mockServer.URL)
	response, err := client.GetDateTime()
	fmt.Printf("%s\n", response)
	fmt.Println(err)

	t.Run("testing successs of response", func(t *testing.T) {
	})
	t.Run("testing correct date", func(t *testing.T) {
		got := strings.Split(string(response), " ")[0]
		want := strings.Split(time.Now().String(), " ")[0]
		assert.Equal(t, want, got)
	})

	t.Run("testing correct time", func(t *testing.T) {
		got := strings.Split(strings.Split(string(response), " ")[1], ":")
		want := strings.Split(strings.Split(time.Now().String(), " ")[1], ":")
		assert.Equal(t, want[:2], got[:2])

		secondsGot, err := strconv.ParseFloat(got[2], 64)
		if err != nil {
			t.Error()
		}
		secondsWant, e := strconv.ParseFloat(want[2], 64)
		if e != nil {
			t.Error()
		}
		if int(secondsWant) != int(secondsGot) && int(secondsWant) != int(secondsGot)+1 {
			t.Error()
		}
	})
	t.Run("testing error return bec of error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, time.Now())
		}))
		defer mockServer.Close()
		client := NewDateTimeClient(mockServer.URL + "blahblah")
		_, err := client.GetDateTime()
		fmt.Println(err)
		assert.NotEqual(t, nil, err)
	})
}

func Test_AllPossibilities(t *testing.T) {
	for i := 0; i < 10; i++ {
		Test_GetDateTime(t)
	}
}
