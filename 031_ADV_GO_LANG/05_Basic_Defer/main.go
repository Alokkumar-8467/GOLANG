package main

func main() {

	// defer resp.body.close()

		fmt.Println("Case 1: success")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Case 2: Fail Early")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
	

}
