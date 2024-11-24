package main

// this is how include multiple packages
import (
	"os"
	"testing"
)

// To make a test create a new file ending in _test.go
// Name of the func should start with Test...   It is mandatory for it to run

// In Visual Studio to run all the commands in the package click on run package tests.
// In Visual Studio to run only the functions in the current file click on file tests.

// to run all tests in a package, run the command => go test

// how to decide what to test => what I care about. What do you think other might break, t he core logic test ?

// to check whether you have covered everything for testing (not very useful) use 
	// go test -cover
// to find out all the lines missed (in text format) use (does it for all files)
	// go test -coverprofile report
// to get above in nice html representation use 
	// go tool cover -html report

// Can run these test cases in parallel (using a configuration) (in case of large number of cases)
// Can use a table driver approach to provide the test cases
// Can use subtests (for proper organisation)
// Writing test cases encourage us to write modular code instead of large fat functions

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 16 {
		t.Errorf("Expected length to be 16, got %v", len(cards))
	}
}

func TestLoadFile(t *testing.T) {
	os.Remove("_decktesting")

	cards := newDeck()
	cards.saveToLocal("_deckTesting")

	newCards := getFromLocal("_deckTesting")
	os.Remove("_decktesting")

	if len(newCards) != 16 {
		t.Errorf("expected 16 cards got %v", len(cards))
	}
}
