package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

// Dictionary struct
// db est un champ prive, qui nous donnera acces a la connection a la base de donnees
type Dictionary struct {
	db *badger.DB
}

// Entry struct
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

// Formatage du created at
func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)

	// 10 nombres de characteres on veut definir
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// Open connection to Database
func New(dir string) (*Dictionary, error) {

	opts := badger.DefaultOptions("./badger")
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)

	if err != nil {
		return nil, err
	}

	dict := &Dictionary{
		db: db,
	}

	return dict, nil
}

// Close connection to database
func (d *Dictionary) Close() {
	d.db.Close()
}
