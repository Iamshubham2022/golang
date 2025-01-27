package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		}
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}
// Get request to get information from the API
func Getreuestring() {

	// Create a custom HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new request
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request using the custom client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	/*
		// Read the response body
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		// fmt.Println("Response (GET):", string(data))

		// Unmarshal the JSON data into a User struct
		var user User
		err = json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println("Error unmarshalling user:", err)
			return
		}
			fmt.Printf("User: %+v\n", user)
	*/

	// the above two functions can be replaced with the single function
	var user User
	error := json.NewDecoder(res.Body).Decode(&user)
	if error != nil {
		fmt.Println("Error unmarshalling user:", error)
		return
	}
	fmt.Printf("User: %+v\n", user)
}

//post method for updation of API
func PostMethod() {
	todo := Todo{
		UserId:    2,
		Id:        1,
		Title:     "shubh jai",
		Completed: true,
	}

	// Convert Todo struct to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling todo:", err)
		return
	}
	fmt.Println("JSON data is:", string(jsonData))

	// API URL for POST request
	myUrl := "https://jsonplaceholder.typicode.com/todos"

	// Convert JSON to string and create reader
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	// Create HTTP client with increased timeout
	client := &http.Client{Timeout: 30 * time.Second}

	// Send POST request
	res, err := client.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer res.Body.Close()

	// Read response body
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print response details
	fmt.Println("Response status:", res.Status)

	fmt.Println("Response body:", string(bodyData))

}

//update the status by using PUT method
func updateStatus(){
	todo := Todo{
		UserId:    2678,
		Id:        1,
		Title:     "shubh jai",
		Completed: true,
	}

	// Convert Todo struct to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling todo:", err)
		return
	}
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req,err:=http.NewRequest("PUT", myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
        return
    }
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	res,err:=client.Do(req)
	if err!= nil {
        fmt.Println("Error sending PUT request:", err)
        return
    }
	defer res.Body.Close()

	bodyData,err:=io.ReadAll(res.Body)
	if err!= nil {
        fmt.Println("Error reading response:", err)
        return
    }
	fmt.Println("Response status:", res.Status)
	fmt.Println("Response body:", string(bodyData))

}

func deleteStatus(){
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req,err:=http.NewRequest("DELETE", myUrl, nil)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
        return
    }

	client := &http.Client{Timeout: 10 * time.Second}
	res,err:=client.Do(req)
	if err!= nil {
        fmt.Println("Error sending DELETE request:", err)
        return
    }
	defer res.Body.Close()
	bodyData,err:=io.ReadAll(res.Body)
	if err!= nil {
        fmt.Println("Error reading response:", err)
        return
    }
	fmt.Println("Response status:", res.Status)
	fmt.Println("Response body:", string(bodyData))
}

func main() {
	fmt.Println("Learning CRUD operations through URL...")
	// Getreuestring()
	// PostMethod()
	// updateStatus()
	// deleteStatus()
}

//How TLS handle

// TLS handshake timeout...........................
/*
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat float32 `json:"lat"`
			Lng float32 `json:"lng"`
		}
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {
	fmt.Println("learning crud through the url...")

	// create, read, update, delete operations
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Error while fetching data from the API:", err)
		return
	}
	defer res.Body.Close()

	//for checking that this url is valid or not
	if res.StatusCode != http.StatusOK {
	    fmt.Printf("Error: %s\n", http.StatusText(res.StatusCode))
		return
	}



	data, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response (GET):", string(data))

	// create
	var user User
	errs := json.Unmarshal(data, &user)
	if errs != nil {
		fmt.Println("Error unmarshalling user:", err)
		return
	}
	fmt.Println(user)
}
*/
