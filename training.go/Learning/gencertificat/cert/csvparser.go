package cert

import (
	"encoding/csv"
	"os"
)

// CSV Parser
func ParseCSV(filename string) ([]*Cert, error) {
	certs := make([]*Cert, 0)

	f, err := os.Open(filename)

	if err != nil {
		return certs, err
	}

	defer f.Close()

	// parser du csv
	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return certs, err
	}

	for _, rec := range records {

		course := rec[0]
		name := rec[1]
		date := rec[2]

		c, err := New(course, name, date)

		if err != nil {
			return certs, err
		}

		// ajout de notre certificat a notre slice (liste)
		certs = append(certs, c)
	}

	return certs, nil
}
