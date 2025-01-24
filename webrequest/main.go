// GET Request: API se data fetch karne ke liye. Example: Fetching a user with ID 1.
// POST Request: New data create karne ke liye. Example: New user create karna.
// PUT Request: Existing data ko update karne ke liye. Example: User details ko update karna.
// DELETE Request: Data ko delete karne ke liye. Example: User ko delete karna.
// PATCH Request: Data ka kuch part update karne ke liye. Example: User ka email address update karna.
// Query Parameters: API mein parameters pass karne ke liye. Example: Filtered user data fetch karna.
// Authentication: API requests mein authentication add karna. Example: API key ya OAuth token ke saath request bhejna.
// JSON Response Handling: API se jo JSON response milta hai, usse parse karna aur data extract karna










// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	fmt.Println("Learning web requests or API handling")

// 	// Sending a GET request to the API
// 	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
// 	if err != nil {
// 		fmt.Println("Error while fetching data from the API:", err)
// 		return
// 	}
// 	defer res.Body.Close() // Break the connection between this code and the website

// 	// Print the type of the response
// 	fmt.Printf("Type of response is: %T\n", res)

// 	// Reading the body of the response
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error while converting data from response:", err)
// 		return
// 	}

// 	// Print the response data as a string
// 	fmt.Println(string(data))
// }







package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

// User structure to unmarshal the response
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// ------------------------ GET Request (Data Retrieval) --------------------------
	// Fetching data from an API using GET
	fmt.Println("GET Request: Fetching data...")
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Error while fetching data:", err)
		return
	}
	defer res.Body.Close()

	// Reading and printing the response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response (GET):", string(data))

	// ------------------------ POST Request (Data Creation) --------------------------
	// Creating data using POST
	fmt.Println("\nPOST Request: Creating data...")
	postData := `{"name": "John", "email": "john@example.com"}`
	res, err = http.Post("https://jsonplaceholder.typicode.com/users", "application/json", bytes.NewBuffer([]byte(postData)))
	if err != nil {
		fmt.Println("Error while posting data:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status (POST):", res.Status)

	// ------------------------ PUT Request (Data Update) ---------------------------
	// Updating data using PUT
	fmt.Println("\nPUT Request: Updating data...")
	putData := `{"name": "John Doe", "email": "john.doe@example.com"}`
	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/users/1", bytes.NewBuffer([]byte(putData)))
	if err != nil {
		fmt.Println("Error while creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while sending PUT request:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status (PUT):", res.Status)

	// ------------------------ DELETE Request (Data Deletion) -----------------------
	// Deleting data using DELETE
	fmt.Println("\nDELETE Request: Deleting data...")
	req, err = http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/users/1", nil)
	if err != nil {
		fmt.Println("Error while creating DELETE request:", err)
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while sending DELETE request:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status (DELETE):", res.Status)

	// ------------------------ PATCH Request (Partial Update) ------------------------
	// Partially updating data using PATCH
	fmt.Println("\nPATCH Request: Partially updating data...")
	patchData := `{"email": "new.email@example.com"}`
	req, err = http.NewRequest("PATCH", "https://jsonplaceholder.typicode.com/users/1", bytes.NewBuffer([]byte(patchData)))
	if err != nil {
		fmt.Println("Error while creating PATCH request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while sending PATCH request:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status (PATCH):", res.Status)

	// ------------------------ Handling Query Parameters -----------------------------
	// Making GET request with query parameters
	fmt.Println("\nGET Request with Query Parameters:")
	url := "https://jsonplaceholder.typicode.com/users"
	req, err = http.NewRequest("GET", url+"?id=1", nil)
	if err != nil {
		fmt.Println("Error while creating GET request:", err)
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while sending GET request:", err)
		return
	}
	defer res.Body.Close()

	// Reading and printing the response
	data, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response with Query Params (GET):", string(data))

	// ------------------------ Handling Authentication -------------------------------
	// Sending request with authentication (API Key)
	fmt.Println("\nGET Request with Authentication (API Key):")
	req, err = http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		fmt.Println("Error while creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer YOUR_API_KEY") // Example of API Key Authentication

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while sending request:", err)
		return
	}
	defer res.Body.Close()

	// Reading and printing the response
	data, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response with Authentication:", string(data))

	// ------------------------ Handling JSON Responses ------------------------------
	// Parsing JSON response to struct
	fmt.Println("\nGET Request: Parsing JSON Response...")
	res, err = http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Error while fetching data:", err)
		return
	}
	defer res.Body.Close()

	var user User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error while decoding JSON:", err)
		return
	}

	fmt.Printf("User Details (JSON Response): ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
}
