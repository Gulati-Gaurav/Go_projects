package main

import "fmt"

func printMap(c map[string]string) {
	// map hence reference and hence can change here its value
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

func main() {
	// ways of declaration

	// 1st way.  [keyType]valueType
	colors := map[string]string{
		"red":    "#ff0000",
		"orange": "#ff1245",
		"pink":   "#ff1233",
	}
	// duplicates not allowed.
	// also commas scenario here as well
	fmt.Println(colors)

	// 2nd way of declaration
	var colors2 map[string]string
	color22 := map[string]string{}
	fmt.Println(color22)

	value, exists := colors2["key"] // to check whether a map contains key or not
	fmt.Println(colors2, value, exists)
	// Unlike C# if we not declare a variable with its value, go initialises with that types' zero value. See the c# code that gives error at bottom

	// 3rd way
	colors3 := make(map[string]string) // equivalent to 2nd. Both creates empty map
	colors3["red"] = "#ff000"          // to put values in map later
	// Both approaches result in an empty map, but map[string]string{} is more common when you might want to initialize the map with some key-value pairs upfront, while make(map[string]string) is more commonly used when the map will be populated later in the code.

	// to delete values
	delete(colors, "red")
	fmt.Println(colors)

	// to iterate over the map
	printMap(colors)

	// to add just do mp["blue"] = "#fd"
	// to access single key value do mp["blue"]
}

// public class HelloWorld
// {
//     Dictionary<string,string> dict;
//     public void name()
//     {
//         Console.WriteLine (dict.ContainsKey("g"));
//     }
//     public static void Main(string[] args)
//     {
//         HelloWorld h = new HelloWorld();
//         h.name();
//     }
// }
