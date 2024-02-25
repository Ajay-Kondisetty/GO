package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type InputData struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

type Address struct {
	ZIP     int    `json:"zip"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type TestInput struct {
	Data   InputData `json:"data"`
	Search string    `json:"search"`
}

func main() {
	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello World!"))
	}
	http.HandleFunc("/", defaultHandler)

	healthCheckHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("I am Alive..."))
	}
	http.HandleFunc("/health_check", healthCheckHandler)

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		requestPayload, _ := io.ReadAll(r.Body)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Printf("some error occurred while reading input: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}(r.Body)
		input := new(TestInput)
		if err := json.Unmarshal(requestPayload, input); err != nil {
			log.Printf("some error occurred while unmarshaling input: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			ref := reflect.ValueOf(input.Data)
			val := ref.FieldByName(strings.Title(input.Search))
			if val.IsValid() {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(fmt.Sprintf("Value of `%s` is : `%v`", input.Search, val)))
			} else {
				innRef := reflect.ValueOf(input.Data.Address)
				innVal := innRef.FieldByName(strings.Title(input.Search))
				if innVal.IsValid() {
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(fmt.Sprintf("Value of `%s` is : `%v`", input.Search, innVal)))
				} else {
					w.WriteHeader(http.StatusNotFound)
					_, _ = w.Write([]byte(fmt.Sprintf("Value of `%s` is not found", input.Search)))
				}
			}
		}
	}
	http.HandleFunc("/test", testHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
