package main

import (
	"fmt"
	"log"

	r "gopkg.in/dancannon/gorethink.v0"
)

func main() {
	fmt.Println("Following Top Scores on RethinkDB")

	session, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.2:28015",
		Database: "players",
	})

	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Table("scores").OrderBy(r.OrderByOpts{Index: r.Desc("Score")}).Limit(5).Changes().Run(session)

	var value interface{}

	if err != nil {
		log.Fatalln(err)
	}
	for res.Next(&value) {
		fmt.Println(value)
	}
}
