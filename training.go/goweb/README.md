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


# REST API OUTIL
``https://insomnia.rest/``
``https://insomnia.rest/download/core/?&ref=`` for Ubuntu


# API SQL de Go (struct DB)

type DB struct {

}

func Open(driverName, dataSourceName string) (*DB, error) 
func OpenDB(c driver.Connector) *DB
func (db *BD) Begin() (*Tx, error)
func (db *DB) Close() error
func (db *DB) Conn(ctx contex.Context) (*Conn, error)
func (db *DB) Driver() driver.Driver
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
func (db *DB) Ping() error
func (db *DB) Prepare(query string) (*Stmt, error)
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
func (db *DB) Stats() DBStats


le package SQL est une interface generique qui doit etre utilisee
avec un pilote SQL

pour ouvrir une base de donnees on va toujours ouvrir notre BD
avec un pilote specifique
Et celui-ci va etre different en fonction de la BD qu'on utilise

Ce pilote est specifique a la BDD que vous utilisez 
 SQLite, MySQL, PostgreSQL

 SQLite sera utilisee dans notre example

 Pour explorer votre BDD, je vous recommande "DB4S"
 voir de facon visuelle notre base de donnees

 Vous pouvez utiliser une autre BDD sans aucun probleme


 # IMPORT DRIVER PACKAGE
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
 ``go get``