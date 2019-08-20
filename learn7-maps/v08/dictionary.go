package dictionary

// Dictionary struct
type Dictionary map[string]string

// DictionaryErr error
type DictionaryErr string

const (
	// ErrNotFound error
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists error
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

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

// Update func
func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}

func (e DictionaryErr) Error() string {
	return string(e)
}
