package dictionary

import (
	"bytes"
	"encoding/gob"
	"sort"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
)

// Add word to the dictionary
func (d *Dictionary) Add(word string, definition string) error {
	entry := Entry{
		Word:       strings.Title(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}

	// convertion d'un struct en bytes (Encodeur / Decodeur binaire)
	// en ecriture on utilise Encode(), en lecture Decode()
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(entry)

	if err != nil {
		return err
	}

	// Castrer (string) en bytes: []byte(word)
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), buffer.Bytes())
	})
}

// Recuperer un mot
func (d *Dictionary) Get(word string) (Entry, error) {

	var entry Entry

	err := d.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(word))

		if err != nil {
			return err
		}

		entry, err = getEntry(item)

		return err
	})

	return entry, err
}

// List retrieves all the dictionary content
// []string is an alphabetically sorted array with the words
// [string]Entry is a map of the words and their definition
func (d *Dictionary) List() ([]string, map[string]Entry, error) {

	entries := make(map[string]Entry)

	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10      // prendra des groupes 10 par 10
		it := txn.NewIterator(opts) // iterateur

		defer it.Close()

		// Rewind - veut dire remettre au debut
		// Valid (tant que valid)
		// Next (faire avancer d'un pas)
		for it.Rewind(); it.Valid(); it.Next() {

			item := it.Item()
			entry, err := getEntry(item)
			if err != nil {
				return err
			}
			entries[entry.Word] = entry
		}

		return nil
	})

	return sortedKeys(entries), entries, err
}

// Sorted Keys
func sortedKeys(entries map[string]Entry) []string {
	keys := make([]string, len(entries))

	for key := range entries {
		keys = append(keys, key)
	}

	// ordonner les cles par ordre alphabetique
	sort.Strings(keys)
	return keys
}

// Get entry
func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry
	var buffer bytes.Buffer

	err := item.Value(func(val []byte) error {
		_, err := buffer.Write(val)

		return err
	})

	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&entry)

	return entry, err
}

// Method for remove the word from the dictionnary
func (d *Dictionary) Remove(word string) error {
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(word))
	})
}
