package sql_request

import (
	"airfilgth/internal/utils"
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Route struct {
	id      int
	origin  string
	arrival string
}

func AddRoute(origin string, arrival string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Route`(origin`, `arrival`) VALUES (?, ?)",
		origin, arrival)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetRoute(selector string, filter string) [][]string {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM Route "
	if filter != "" {
		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val [][]string
	var tag Route
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.origin, &tag.arrival)
		to_inject := []string{strconv.Itoa(tag.id), tag.origin, tag.arrival}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdateRoute(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `Route` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteRoute(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `Route` WHERE " + condition

	db.Query(query)

}
