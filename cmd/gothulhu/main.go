package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ropfoo/gothulhu/internal/character"
)

var reader = bufio.NewReader(os.Stdin)

var writer = bufio.NewWriter(os.Stdout)

func main() {
	options := getOptions()
	fmt.Println(options)

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Age: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	name = strings.TrimSpace(name)
	ageInt, _ := strconv.Atoi(age)

	character := character.GenerateCharacter(name, ageInt)
	if options.ShouldExportJSON {
		// export to json file
		fileName := fmt.Sprintf("%s.json", name)
		writer.WriteString(character.ToJSON())
		os.WriteFile(fileName, []byte(character.ToJSON()), 0644)
	} else {
		writer.WriteString(character.ToJSON())
	}
	writer.Flush()
}

type Options struct {
	ShouldExportJSON bool
}

// get options from console input
func getOptions() Options {
	// Check if "--json" flag is present in any argument
	for _, arg := range os.Args[1:] {
		if arg == "--json" {
			return Options{
				ShouldExportJSON: true,
			}
		}
	}
	return Options{}
}
