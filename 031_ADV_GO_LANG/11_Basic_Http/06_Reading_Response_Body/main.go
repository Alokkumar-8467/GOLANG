package main

func main() {

}
func main() {

	url := "http://jsonplaceholder.typicode.com/todos"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

}
