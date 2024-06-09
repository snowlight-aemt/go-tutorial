package mydict

import (
	"errors"
)

type Dictionary map[string]string

var errNotFound = errors.New("not Found")

var errAlreadyExists = errors.New("already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return value, errNotFound
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = value
		return nil
	}

	return errAlreadyExists
}
