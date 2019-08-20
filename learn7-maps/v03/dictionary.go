package dictionary

// Dictionary struct
type Dictionary map[string]string

// Search func
func (d Dictionary) Search(word string) string {
	return d[word]
}
