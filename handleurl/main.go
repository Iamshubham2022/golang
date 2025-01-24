package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 1. Parsing a URL
	originalURL := "https://example.com/path/to/resource?name=John&age=30#section1"
	parsedURL, err := url.Parse(originalURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Original URL:", parsedURL.String())

	// 2. Extracting URL components
	fmt.Println("Scheme:", parsedURL.Scheme)      // https
	fmt.Println("Host:", parsedURL.Host)          // example.com
	fmt.Println("Path:", parsedURL.Path)          // /path/to/resource
	fmt.Println("Raw Query:", parsedURL.RawQuery) // name=John&age=30
	fmt.Println("Fragment:", parsedURL.Fragment)  // section1

	// 3. Accessing and modifying query parameters
	queryParams := parsedURL.Query()
	for key, values := range queryParams {
		fmt.Printf("Query Param: %s = %s\n", key, values[0]) // name=John, age=30
	}

	// Modifying query parameters
	queryParams.Set("name", "Alice")
	queryParams.Add("status", "active")
	parsedURL.RawQuery = queryParams.Encode()

	fmt.Println("\nModified Query:", parsedURL.RawQuery)
	fmt.Println("URL after modifying query:", parsedURL.String())

	// 4. Modifying other URL components
	parsedURL.Scheme = "http"           // Changing scheme
	parsedURL.Host = "newdomain.com"    // Changing host
	parsedURL.Path = "/new/path/to/api" // Changing path
	parsedURL.Fragment = "updatedSection"

	fmt.Println("\nURL after modifications:", parsedURL.String())

	// 5. Building a new URL from scratch
	newURL := &url.URL{
		Scheme:   "https",
		Host:     "api.example.com",
		Path:     "/v1/resources",
		RawQuery: "id=123&status=active",
		Fragment: "details",
	}
	fmt.Println("\nNewly Built URL:", newURL.String())

	// 6. Encoding a URL
	baseURL := "https://example.com/search"
	params := url.Values{}
	params.Add("query", "Go programming")
	params.Add("page", "1")
	encodedURL := baseURL + "?" + params.Encode()
	fmt.Println("\nEncoded URL with parameters:", encodedURL)

	// 7. Decoding a query string
	decodedParams, _ := url.ParseQuery("query=Go+programming&page=1")
	fmt.Println("\nDecoded Query Parameters:")
	for key, value := range decodedParams {
		fmt.Printf("%s = %s\n", key, value[0]) // query=Go programming, page=1
	}
}



// Original URL: https://example.com/path/to/resource?name=John&age=30#section1
// Scheme: https
// Host: example.com
// Path: /path/to/resource
// Raw Query: name=John&age=30
// Fragment: section1
// Query Param: name = John
// Query Param: age = 30

// Modified Query: age=30&name=Alice&status=active
// URL after modifying query: https://example.com/path/to/resource?age=30&name=Alice&status=active

// URL after modifications: http://newdomain.com/new/path/to/api?age=30&name=Alice&status=active#updatedSection

// Newly Built URL: https://api.example.com/v1/resources?id=123&status=active#details

// Encoded URL with parameters: https://example.com/search?query=Go+programming&page=1

// Decoded Query Parameters:
// query = Go programming
// page = 1
