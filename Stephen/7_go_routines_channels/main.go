package main

func main() {

	urlList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// withoutRoutines(urlList)
	// withGoRoutine(urlList)
	// withGoRoutineChannel(urlList)
	withContinously(urlList)
}
