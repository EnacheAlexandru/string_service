package main

import (
	"bufio"
	"os"
	"strings"
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

	for scanner.Scan() {
		var line string = scanner.Text()

		// I assumed from the task description that the header always contains lowercase letters
		// I also assumed that the first letter of each line (header excluded) is an uppercase letter
		// because the first field is supposed to be the name of a person
		if strings.ToLower(line) == line {
			continue
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

	// complexity of O(n) - we are only iterating once through the file and through the keys/values of the dictionary
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
	Transformation("input.csv", "output.csv")
}
