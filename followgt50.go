package main

import (
	"fmt"
	"log"

	r "gopkg.in/dancannon/gorethink.v0"
)

func main() {
	fmt.Println("Following Greater than 50 on RethinkDB")

	session, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.2:28015",
		Database: "players",
	})

	if err != nil {
		log.Fatal(err)
	}

	res, err := r.Table("scores").Filter(r.Row.Field("Score").Gt(50)).Changes().Run(session)
	//res, err := r.Table("scores").Filter(r.Row.Field("Score").Gt(50)).Pluck("PlayerName").Changes().Run(session)

	var value interface{}

	if err != nil {
		log.Fatalln(err)
	}
	for res.Next(&value) {
		fmt.Println(value)
	}
}
