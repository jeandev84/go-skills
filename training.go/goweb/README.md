# GO AND WEB
``SERVER HTTP``

package main

import (
	"log"
	"net/http"
)

func main() {

	/*
		err := http.ListenAndServe(":8000", nil)
		if err != nil {

		}

		revient a :
		if err := http.ListenAndServe(":8000", nil); err != nil {

		}
	*/

	// Lancer le server en ecoutant au port 8000
	// go build && ./goweb
	// port :80 (site non securise http), port :443 (site securise https)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

    server(8000)
}

func server(port string) {

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}
