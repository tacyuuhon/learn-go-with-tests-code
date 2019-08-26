package dictionary

import (
	"errors"
)

// Dictionary struct
type Dictionary map[string]string

// ErrNotFound error
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search func
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add func
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
