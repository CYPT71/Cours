package sql

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct {
	id         int
	ticket_id  int
	bank       string
	profession string
	name       string
	first_name string
	adress     string
}

func AddPassenger(profession string, ticket_id int, bank int, name string, first_name string, address string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Passenger`(`name`, `first_name`, `adress`, `profession`, `bank`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?)",
		name, first_name, address, profession, bank, ticket_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetPassenger(selector string, filter string) [][]string {

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
	query += "FROM Passenger "
	if filter != "" {
		query += " WHERE `id` IN (" + filter + ")"
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val [][]string
	var tag Passenger
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.adress, &tag.bank, &tag.profession, &tag.ticket_id, &tag.first_name)
		to_inject := []string{strconv.Itoa(tag.id), tag.adress, strconv.Itoa(tag.ticket_id), tag.bank,
			tag.name, tag.first_name}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdatePassenger(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "UPDATE `Passenger` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeletePassenger(condition string) {
	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "DELETE FROM `Passenger` WHERE " + condition

	db.Query(query)

}
