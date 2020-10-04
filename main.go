package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// IsValidVerb - Checks if the noun is a valid param
func IsValidVerb(verb string) bool {
	switch verb {
	case
		"add",
		"latest":
		return true
	}
	return false
}

// Error - Message returned to user
type Error struct {
	message string
}

// Tag - Holds the tag being discovered
type Tag struct {
	noun string
	verb string
	file string
	tag  *NewTag
	err  *Error
}

// NewTag - For adding new tags to file
type NewTag struct {
	tag string
}

// Found - resulting tag
type Found struct {
	tag string
}

func (t *Tag) verify() (bool, *Error) {

	if t.noun != "tag" {
		t.err.message = t.noun + " not supported"
		return false, t.err
	}

	if IsValidVerb(t.verb) != true {
		t.err.message = "Needs to be: tag, add, latest or delete"
		return false, t.err
	}

	if t.file == "" {
		t.err.message = "No input file provided"
		return false, t.err
	}

	return true, nil
}

func (t *Tag) execute() (Found, *Error) {

	found := Found{}

	switch t.verb {
	case "latest":
		tag, err := getLatestTag(t)

		if err != nil {
			t.err.message = "Could not find tag"
			panic("Something went wrong finding the latest tag")
		}

		found.tag = tag

		break
	case "add":
		if t.tag.tag == "" {
			panic("No tag provided could not add")
		}

		tag, err := addLineToFile(t.file, t.tag.tag)

		if err != nil {
			t.err.message = "Could not add tag"
			panic("Something went wrong adding the latest tag")
		}

		found.tag = tag

	}

	if found.tag != "" {
		return found, nil
	}

	return found, t.err
}

func main() {

	var filename string

	if len(os.Args) < 3 {
		panic("tagfinder expected at least 2 arguments to be passed")
	}

	flag.StringVar(&filename, "fpath", "tags.txt", "file path to read from")
	flag.Parse()

	args := flag.Args()

	// fmt.Println(args)

	if filename == "" {
		panic("No file provided")
	}

	var newTag NewTag = NewTag{}

	if len(args) > 2 {
		newTag.tag = args[2]
	}

	tag := Tag{args[0], args[1], filename, &newTag, &Error{}}

	// Verify arguments are as expected
	_, err := tag.verify()

	if err != nil {
		panic("Error with request" + err.message)
	}

	found, err := tag.execute()

	if err != nil {
		log.Fatal("Tag could not be found")
	}

	fmt.Println(found.tag)
}
