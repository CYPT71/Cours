package sql

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	id       int
	capacity int
	types    string
}

func AddDevices(capacity int, types string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Device` (`capacity`, `type`) VALUES (? , ?)", capacity, types)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetDevices(selector string, filter string) [][]string {

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

	var return_val [][]string
	var tag Device
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.capacity, &tag.types)
		to_inject := []string{strconv.Itoa(tag.id), strconv.Itoa(tag.capacity), tag.types}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdateDevice(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "UPDATE `Device` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteDevice(condition string) {
	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "DELETE FROM `Device` WHERE " + condition

	db.Query(query)

}
