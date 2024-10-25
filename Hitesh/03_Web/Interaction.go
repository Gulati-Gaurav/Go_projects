package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const myServerUrl = "http://localhost:8000"

// Make sure you keep your server up and running

func MakeGetRequest(param string) {

	resp, err := http.Get(myServerUrl + param)
	CheckErrNil(err)
	defer resp.Body.Close()

	fmt.Println("Status Code: ", resp.Status)

	// io.Copy(os.Stdout, resp.Body)
	// or

	var respString strings.Builder
	content, err := io.ReadAll(resp.Body)
	byteCount, err := respString.Write(content)
	fmt.Println(byteCount, respString.String()) // to convert builder to string
}

func MakePostRequest(param string) {
	// hack to make sample json data
	content := `
		{
			"Message": "let's go with golang",
			"Price": 12,
			"Platform": "LCO.in"
		}
	`
	// In quotes "" you need to escape new lines, tabs and other characters that do not need to be escaped in backticks ``.
	// we use `` for multi line string

	// Now how to pass something which has our content and also implements the reader interface ?
	reqBody := strings.NewReader(content)
	// or
	// reqBody := bytes.NewBufferString(content)

	resp, err := http.Post(myServerUrl+param, "application/json", reqBody)
	CheckErrNil(err)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func MakeFormPostRequest(param string) {
	// in case of image upload form data is necessary
	data := url.Values{}
	data.Add("firstName", "hitesh")
	data.Add("lastName", "Gulati")
	data.Add("School", "BBPS")

	resp, err := http.PostForm(myServerUrl+param, data)
	CheckErrNil(err)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

// make sure to name the fields capital since this struct will be called in below method where it is not a receiver func
type course struct {
	// make sure to have no spaces in the json `` below
	Name     string `json:"courseName"` // incase you wanna store the key in json of different name
	Price    int
	Platform string
	Password string   `json:"-"`              // means we don't wanna show this in response
	Tags     []string `json:"tags,omitempty"` // means if value is nil don't send at all
}

func encodeJson() {
	lcoCourse := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}
	// package as json data
	// Marshal returns the JSON encoding of the arguement.
	// finalJson, err := json.Marshal(lcoCourse)

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	CheckErrNil(err)
	fmt.Printf("%s\n", finalJson)
	// or
	// fmt.Println(string(finalJson))

}

func decodeJson() {
	// we always receive data from internet in byte format
	jsonData := []byte(`
	{	
		"Name": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Password": "abc123",
		"Tags": ["web-dev","js"]
	}
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonData) // to check what we received was valid in the first place or not
	if checkValid {
		fmt.Println("JSON was valid")
		err := json.Unmarshal(jsonData, &lcoCourse) // & so that copy is not created
		CheckErrNil(err)
		fmt.Printf("%#v", lcoCourse)
	}

	// but if you don't wanna cast it into a class but play with key value
	var myOnlineData map[string]any // since we don't always know value type
	err := json.Unmarshal(jsonData, &myOnlineData)
	CheckErrNil(err)
	fmt.Print(myOnlineData) // can loop through this as well
}
