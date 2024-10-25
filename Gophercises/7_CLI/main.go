package main

import (
	"cliapp/cmd"
)

func main() {
	cmd.Execute()
}

// Later you can dig into how to install the application so that it can be run from anywhere and it will persist our tasks regardless of where we run the CLI.

// After that, explore how to setup and install the application so that it can be run from any directory in your terminal. This might require you to look into how to find a user's home directory on any OS (Windows, Mac OS, Linux, etc).

// If you'd like, you can look into how to determine this on your own, but I recommend just grabbing this package: https://github.com/mitchellh/go-homedir. You can read over the code to see how it works - it is only 137 lines of code - but it should take care of all the oddities between different operating systems for us.

// After that you will need to look into how to install a binary on your computer. The first place I suggest starting is the go install command. (Hint: Try go install --help to see what this command does.). This is likely to be the simplest route, but there are other options (like manually copying a binary to a directory in your $PATH).

// Note: I suspect many users will have issues around here that are OS specific. If you do, please first check the Github issues to see if there are any open or closed issues that are similar to your problem. I'm hoping to use that as a nice Q&A section for this exercise.
