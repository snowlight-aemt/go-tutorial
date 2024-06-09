package mydict

import (
	"errors"
)

type Dictionary map[string]string

var (
	errNotFound      = errors.New("not Found")
	errCanNotUpdate  = errors.New("can not update no existing word")
	errAlreadyExists = errors.New("already exists")
)

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

func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCanNotUpdate
	}

	d[word] = value
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return err
	}

	delete(d, word)
	return nil
}
