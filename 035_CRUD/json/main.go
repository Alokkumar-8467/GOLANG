	fmt.Println("Person Data after marshaling is : ", string(jsonData))

	// converting JSON Decoding (Unmarshal)
	// Means converting JSON data into struct

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Unmarshaled Person Data is : ", personData)
}

// So Direct Structure Data is :  {John Doe 30 true}
// Person Data after marshaling is :  {"name":"John Doe","age":30,"isAdult":true}
// Unmarshaled Person Data is :  {John Doe 30 true}
