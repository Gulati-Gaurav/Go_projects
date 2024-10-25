package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myUrl2 = "https://www.google.co.in/"

func Web() {
	// web requests through http/web are never closed itself, we have to close them ourselves else if they redirect to any other website - you are responsible.
	resp, err := http.Get(myUrl2)
	CheckErrNil(err)

	defer resp.Body.Close() // your resposibility to close the connection

	data, err := io.ReadAll(resp.Body)
	CheckErrNil(err)
	fmt.Println(string(data))
}

func CheckErrNil(err error) {
	if err != nil {
		panic(err)
	}
}

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb#category"

func Urls() {
	// like http in .net we have url in go
	// url package
	// allows you to parse URLs into their components, such as scheme, host, path, and query parameters
	// and provides methods to escape and unescape query parameters, ensuring that they are correctly encoded for safe transmission over the internet.

	// The general form represented is:
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	// https: The protocol used (Hypertext Transfer Protocol Secure)
	// user:password: The username and password for authentication (userinfo)
	// www.example.com: The domain name of the server (host)
	// path/to/file.html: The location of the file on the server
	// param1=value1&param2=value2: Query parameters (optional)
	// section: Specifies a particular part of the resource. It's often used for direct linking to specific content within a webpage, such as a section heading, image, or interactive element. For instance, in a long article, you might use a fragment identifier to link directly to the conclusion.

	resp, err := url.Parse(myUrl)
	CheckErrNil(err)

	fmt.Println("Scheme : ", resp.Scheme)
	fmt.Println("Host : ", resp.Host)
	fmt.Println("Port : ", resp.Port())
	fmt.Println("Path : ", resp.Path)
	fmt.Println("RawQuery : ", resp.RawQuery)
	fmt.Println("Fragment : ", resp.Fragment)

	// Access and manipulate query parameters.
	//RawQuery just returns a string. Query() returns map after iterating over RawQuery
	queryParams := resp.Query()
	fmt.Println("Query Parameters:", queryParams)
	fmt.Println(queryParams["coursename"])

	queryParams.Add("page", "1")         // Add a new query parameter in alphabetical order
	resp.RawQuery = queryParams.Encode() // Encode the query parameters back into the URL since we can't assign resp.Query() to anything

	fmt.Println(resp) // with updated queryParams

	partsOfUrl := &url.URL{ // builds a url for us. Remember the &
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=hitesh",
	}
	fmt.Println(partsOfUrl.String())

	// Efficiency: Pointers can be more efficient than directly copying large objects because you're only passing the memory address, not the entire object. This is especially beneficial when dealing with complex structures like URLs.
	// Modification: Since partsOfUrl points to the actual url.URL object in memory, any changes made through partsOfUrl will directly affect the original object.

	// Hence whenever receiving a struct from any function try to receive its pointer not the struct itself
}
