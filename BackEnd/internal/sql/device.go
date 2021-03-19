package sql

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	id       int    `json:"id"`
	capacity int    `json:"capacity"`
	types    string `json:"type"`
}

func AddDevices(capacity int, types string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	fmt.Printf("INSERT INTO Departus (type, capacity) VALUES (" + strconv.Itoa(capacity) + "," + types + ")")
	insert, err := db.Query("INSERT INTO Departus (type, capacity) VALUES (" + strconv.Itoa(capacity) + "," + types + ")")

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetDevice(selector string, filter string) [][]string {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM Departus "
	if filter != "" {
		query += " WHERE `id` IN (" + filter + ")"
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var departures [][]string
	var tag Device
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.capacity, &tag.types)
		to_inject := []string{strconv.Itoa(tag.id), strconv.Itoa(tag.capacity), tag.types}
		departures = append(departures, to_inject)
	}

	return departures

}
