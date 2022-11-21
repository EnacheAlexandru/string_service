package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// complexity of O(n)
func Transformation(ifilename string, ofilename string) {

	finput, ferror := os.Open(ifilename)

	// check if the input file exists or opens properly
	if ferror != nil {
		panic("Input file not found or failed to open!")
	}

	defer finput.Close()

	// maps have been used because they have a complexity of O(1) when searching for a key

	// Key: a line from the input file (header excluded)
	// Value: irrelevant; we are only interested to see if a certain line exists in order to avoid duplicates
	noDuplicates := map[string]bool{}

	// Key: the first letter that groups the lines
	// Value: a string that has all the lines concatened, grouped by a certain first letter (e.g. A: Anita...\nAron...\n)
	letterToGroup := map[string]string{}

	scanner := bufio.NewScanner(finput)
	isHeaderRead := false

	//complexity of O(n) - we are only iterating once through the file
	for scanner.Scan() {
		var line string = scanner.Text()

		// The header should have the first field "full_name". The fields are lowercase, separated with comma and space.
		if !isHeaderRead {
			reg, _ := regexp.Compile("^full_name(, [a-z_]+)*$")
			if reg.MatchString(line) {
				isHeaderRead = true
				continue
			}
			panic("The header should have the first field \"full_name\". The fields are lowercase, separated with comma and space.")
		}

		// The first letter of each line (header excluded) is an uppercase letter because the first field is supposed to be the name of a person
		reg, _ := regexp.Compile("^[A-Z]$")
		if !reg.MatchString(line[0:1]) {
			panic("The first letter of each line (header excluded) should be an uppercase letter!")
		}

		_, lineExists := noDuplicates[line]

		// if there is no duplicated line (the key is not already present in the map)
		if !lineExists {

			// we add the line
			noDuplicates[line] = true

			// grabs the first character from a string
			var letter string = line[0:1]

			_, letterExists := letterToGroup[letter]

			// if the letter does not exist in the map
			if !letterExists {
				letterToGroup[letter] = line + "\n"
			} else {
				letterToGroup[letter] += line + "\n"
			}
		}

	}

	var output string = ""

	// complexity of O(k) - we are only iterating once through the pair of keys/values (k < n)
	// bulding the output string
	for letter, group := range letterToGroup {
		output += letter + ":\n" + group + "\n"
	}

	foutput, ferror := os.Create(ofilename)

	// check if there are errors when creating the output file
	if ferror != nil {
		panic("Error creating output file!")
	}

	defer foutput.Close()

	// the output string will always finish with \n\n, because WriteString adds a \n
	// we want to remove one of the \n
	if len(output) > 0 {
		output = output[:len(output)-1]
	}

	// =!!!= because maps are unordered, each run generates a random order for the groups =!!!=
	_, ferror2 := foutput.WriteString(output)

	// check if there are errors when writing to output file
	if ferror2 != nil {
		panic("Error writing to output file!")
	}

}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred:", err)
		}
	}()

	Transformation("input.csv", "output.csv")
}
