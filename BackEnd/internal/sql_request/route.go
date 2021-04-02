package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Route struct {
	Id      int
	Origin  string
	Arrival string
}

func AddRoute(Origin string, Arrival string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `route`(`origin`, `arrival`) VALUES (?, ?)",
		Origin, Arrival)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetRoute(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `route` "
	if filter != "" {
		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	var tag Route
	for selecte.Next() {
		selecte.Scan(&tag.Id, &tag.Origin, &tag.Arrival)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Id":      tag.Id,
			"Origin":  tag.Origin,
			"Arrival": tag.Arrival,
		})
	}

	return return_val

}

func UpdateRoute(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `route` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteRoute(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `route` WHERE " + condition

	db.Query(query)

}
