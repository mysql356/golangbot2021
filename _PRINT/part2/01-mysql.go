﻿package main

import (
    "log"
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

/*
 * Tag... - a very simple struct
 */
type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func mysql() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

	//Insert/////////////////////////////////////////////////
    // perform a db.Query insert
    insert, err := db.Query("INSERT INTO test VALUES ( null, now() )")

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()

	//one row ///////////////////////////////////////////
	var tag Tag
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM test where id = ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(tag.ID)
	log.Println(tag.Name)

	//multiple row ///////////////////////////////////////////////////
    // Execute the query
    results, err := db.Query("SELECT id, name FROM test")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Println(tag.ID, " ============> ", tag.Name)
    }


}