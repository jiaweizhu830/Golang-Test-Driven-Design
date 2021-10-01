package main

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

// make error constant by using DictionaryErr type
// https://dave.cheney.net/2016/04/07/constant-errors
var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Customized dictionary
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// ok indicates if the key was found successfully
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// Go's built-in function (delete for map)
	delete(d, word)
}

/**
Map: can modify them WITHOUT passing as an address to it!
When you pass a map to a func, you are indeed copying it, but just the pointer part

Map can be a nil.
A nil map behaves like an empty map when reading.
When writing to a nil map, will cause a runtime panic
So, should NEVER initialize an empty map variable like: var m map[string]string

Instead, initialize an empty map like below (will create an empty hash map & point dictionary at it)
	ensure never get a runtime panic
1. var dictionary = map[string]string{}
2. var dictionary = make(map[string]string)
*/
