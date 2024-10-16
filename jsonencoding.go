package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func jsonencoding() {

	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		todoitem := Todo{}

		json.Unmarshal(bodyBytes, &todoitem)

		fmt.Println(todoitem)
	}

}
