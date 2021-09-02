package main

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func dbConn() (db *sql.DB) {
        db, err := sql.Open("sqlite3", "./bc.db")
        // Check if database connection was opened successfully
        if err != nil {
                // Print error and exit if there was problem opening connection.
                log.Fatal(err)
        }
        // close database connection before exiting program.
        return db
}
