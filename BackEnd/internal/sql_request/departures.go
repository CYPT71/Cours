package sql_request

import (
	"airfilgth/internal/utils"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Departus struct {
	Id          int
	Date        time.Time
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
	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM departures "
	if filter != "" {
		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	var tag Departus
	for selecte.Next() {
		selecte.Scan(&tag.Id, &tag.Date, &tag.Pilote, &tag.Copilote, &tag.Aircrew, &tag.Free_places, &tag.Occupied)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"id":         tag.Id,
			"Pilote":     tag.Pilote,
			"Copilote":   tag.Copilote,
			"Aircrew":    tag.Aircrew,
			"Free Place": tag.Free_places,
			"Occupied":   tag.Occupied,
		})
	}

	return return_val

}

func AddDepartures(id_flight int, date time.Time, pilote int, copilote int, aircrew string, free_places int, occupied int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `departures`(`id_fligth`, `date`, `pilote`, `copilote`, `aircrew`, `free_places`, `occupied`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		id_flight, date, pilote, copilote, aircrew, free_places, occupied)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func UpdateDepartus(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `departures` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteDepartus(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	query := "DELETE FROM `departures` WHERE " + condition

	db.Query(query)

}
