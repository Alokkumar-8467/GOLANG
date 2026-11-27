package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/todos"
	url1 := "http://jsonplaceholder.typicode.com/todos/1"
	url2 := "http://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url)
	resp, err := http.Get(url1)
	resp2, err := http.Get(url2)

	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}

	defer res.Body.Close()
	defer resp.Body.Close()

	// We can also check the status code and status message of the response to see if the request was successful or not. The status code is a three-digit number that indicates the outcome of the request, while the status message is a brief description of the status code.

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response: ", resp.Status)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	fmt.Println("Data : ", string(data))

	fmt.Println("Status code", resp.StatusCode)
	fmt.Println("Status code", resp.Status)

}

/* Output
Status code 200
Status code 200 OK
*/
