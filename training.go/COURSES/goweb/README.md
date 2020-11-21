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

 ``Pour explorer votre BDD, je vous recommande "DB4S"``
 voir de facon visuelle notre base de donnees

 Vous pouvez utiliser une autre BDD sans aucun probleme


 # IMPORT DRIVER PACKAGE
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
 ``go get``


# CREER UNE TABLE SQL
CREATE TABLE IF NOT EXISTS book
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	author TEXT,
	page_count INTEGER
)
``https://github.com/sqlitebrowser/sqlitebrowser``


Ubuntu and Derivatives
Stable release
For Ubuntu and derivaties, @deepsidhu1313 provides a PPA with the latest release here:

https://launchpad.net/~linuxgndu/+archive/ubuntu/sqlitebrowser
To add this ppa just type in these commands in terminal:

sudo add-apt-repository -y ppa:linuxgndu/sqlitebrowser
Then update the cache using:

sudo apt-get update
Install the package using:

sudo apt-get install sqlitebrowser
Ubuntu 14.04.X, 15.04.X, 15.10.X and 16.04.X are supported for now (until Launchpad decides to discontinue building for any series).

Ubuntu Precise (12.04) and Utopic (14.10) are not supported:

Precise does not have a new enough Qt package in its repository by default, which is a dependency
Launchpad does not support Utopic any more, which has reached its End of Life
Nightly builds
Nightly builds are available here:

https://launchpad.net/~linuxgndu/+archive/ubuntu/sqlitebrowser-testing
To add this ppa, type these commands into the terminal:

sudo add-apt-repository -y ppa:linuxgndu/sqlitebrowser-testing
Then update the cache using:

sudo apt-get update
Install the package using:

sudo apt-get install sqlitebrowser
Other Linux


# REQUETTES PREPAREES
db, _ := sql.Open("sqlite3", "test.db")
stmt, _ = db.Prepare("INSERT INTO user (name, email) VALUES(?, ?)")
stmt.Exec("Bob", "bob@golang.org")