package myDict

import "errors"

// Dictionary type
type Diiiii map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

/**
 * Map types are reference types, like pointers or slices.
 * So when we hand it over to the arg or get by receiver, we'll get itself(not a copy).
 * 
 * that's why it doesn't have to receive pointer.
 */

// Search for a word
func (d Diiiii) Search(word string) (string, error) {
	/**
	## i, ok := m["route"]

	In this statement, the first value (i) is assigned the value stored under the key "route".
	If that key doesn’t exist, i is the value type’s zero value (0).
	The second value (ok) is a bool that is true if the key exists in the map, and false if not.

	https://go.dev/blog/maps
	 */
	value, exists := d[word]
	
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Diiiii) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}