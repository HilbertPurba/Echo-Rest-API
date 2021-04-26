package main

import (
	"github.com/hilbertpurba/PBP/tugas-crud-echo/db"
	"github.com/hilbertpurba/PBP/tugas-crud-echo/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
