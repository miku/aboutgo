// Web capture tool. Takes snapshots of website content and stores it into
// sqlite database.
//
// $ sqlite3 data.db < create.sql
//
//
// Exercises (choose one)
//
// 1) Add a new subcommand - e.g. "show" - that takes an ID of a capture  and
// displays the webpage.
//
// 2) Add a new subcommand to show all captures of a given URL.
//
// Try to utilize https://jmoiron.github.io/sqlx/#getAndSelect shortcuts to
// deserialize a row into a WebCapture struct.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// WebCapture saves a URL and its content to a database database.
type WebCapture struct {
	ID         int    `db:"id"`
	Link       string `db:"link"`
	Text       string `db:"text"`
	StatusCode int    `db:"statuscode"`
	Time       int    `db:"time"`
}

func prependSchema(s string) string {
	if strings.HasPrefix(s, "http") {
		return s
	}
	return fmt.Sprintf("http://%s", s)
}

func main() {
	showID := flag.Int("i", 0, "specify the id to show")

	flag.Parse()

	db, err := sqlx.Connect("sqlite3", "data.db") // Open and connect.
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var mode string
	if flag.NArg() > 0 {
		mode = flag.Arg(0)
	}

	var insertStmt = `INSERT INTO webcapture (link, text, statuscode, time) VALUES (?, ?, ?, ?)`

	switch mode {
	case "s", "show":
		var wc WebCapture
		err := db.Get(&wc, "SELECT text FROM webcapture WHERE id = ?", showID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wc.Text)
	case "g", "get":
		if flag.NArg() < 2 {
			log.Fatal("use: get URL")
		}

		link := prependSchema(flag.Arg(1))

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		result, err := db.Exec(insertStmt, link, string(b), resp.StatusCode, time.Now().Unix()) // sql.driverResult
		if err != nil {
			log.Fatal(err)
		}

		lid, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		ra, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("last_insert_id=%v, rows_affected=%v", lid, ra)
	default:
		var wcs []WebCapture

		// Using db.Select - https://jmoiron.github.io/sqlx/#getAndSelect
		err := db.Select(&wcs, "SELECT * FROM webcapture ORDER BY time DESC LIMIT 10")
		if err != nil {
			log.Fatal(err)
		}
		for _, wc := range wcs {
			date := time.Unix(int64(wc.Time), 0)
			fmt.Printf("%v\t%v\t%v\t%v\t%v\n", wc.ID, wc.StatusCode, date, len(wc.Text), wc.Link)
		}
	}
}
