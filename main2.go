/*package main

/* import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  Title  `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

type Title struct {
	MainTitle   string `json:"main_title"`
	SecondTitle string `json:"secondary_title"`
}

func main() {
	title := Title{MainTitle: "Learning Concurrency", SecondTitle: "Learning Conucrrency in Python"}
	author := Author{Name: "Elliot Forbes", Age: 25, Developer: true}
	book := Book{Title: title, Author: author}

	fmt.Printf("%+v\n", book)

	fmt.Println("\n", book, "\n")

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray)) */

/*import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Book struct {
	Title  Title  `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

type Title struct {
	MainTitle   string `json:"main_title"`
	SecondTitle string `json:"secondary_title"`
}

func main() {

	// Open the JSON file
	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after we are done

	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into the struct
	var test Book
	err = json.Unmarshal(data, &test)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the result
	fmt.Println(test)
	fmt.Printf("%+v\n", test)

} */

package main

import (
	"encoding/json"
	"fmt"
	"io/io"
	"os"
)

type Book struct {
	Title  Title  `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

type Title struct {
	MainTitle   string `json:"main_title"`      // Exported
	SecondTitle string `json:"secondary_title"` // Exported
}

func testJson() {
	// Open the JSON file
	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after we are done

	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into the struct
	var test Book
	err = json.Unmarshal(data, &test)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the result
	fmt.Println(test.Title.MainTitle)

	fmt.Printf("%+v\n", test) // Use %+v to see field names and values
}
