package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpMethod := "GET"
	url := "http://localhost:8080"

	client := http.Client{} // use this client to perform any http method you need

	request, err := http.NewRequest(httpMethod, url, nil)

	// configure headers
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body) // Read from response until an error or EOF and return the data it read
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
