package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Departures struct {
	Id          int
	Date        string
	Pilote      int
	Copilote    int
	Aircrew     string
	Free_places int
	Occupied    int
}

func GetDepartures(selector string, filter string) []map[string]interface{} {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM `departures` "
	if filter != "" {
		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag Departures
		selecte.Scan(&tag.Id, &tag.Date, &tag.Pilote, &tag.Copilote, &tag.Aircrew, &tag.Free_places, &tag.Occupied)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Id":         tag.Id,
			"Pilote":     tag.Pilote,
			"Copilote":   tag.Copilote,
			"Aircrew":    tag.Aircrew,
			"Free Place": tag.Free_places,
			"Occupied":   tag.Occupied,
		})
	}

	return return_val

}

func AddDepartures(Id_flight int, Date string, Pilote int, Copilote int, Aircrew string, Free_places int, Occupied int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `departures`(`id_flight`, `date`, `pilote`, `copilote`, `aircrew`, `free_places`, `occupied`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		Id_flight, Date, Pilote, Copilote, Aircrew, Free_places, Occupied)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func UpdateDepartures(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "UPDATE `departures` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteDepartures(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "DELETE FROM `departures` WHERE " + condition

	db.Query(query)

}
