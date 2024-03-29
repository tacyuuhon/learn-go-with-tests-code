package dictionary

import (
	"errors"
)

// Dictionary struct
type Dictionary map[string]string

// ErrNotFound error
var ErrNotFound = errors.New("could not find the word you were looking for")

// ErrWordExists error
var ErrWordExists = errors.New("cannot add word because it already exists")

// Search func
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add func
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
