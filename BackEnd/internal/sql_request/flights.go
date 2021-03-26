package sql_request

import (
	"airflight/internal/utils"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Flight struct {
	Id            int
	Id_departures int
	Arrival       string
	Id_route      int
	Id_device     int
}

func AddFlight(Id_departures int, Arrival string, Id_route int, Id_device int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `flight`(`id_departures`, `arrival`, `id_route`, `id_device`) VALUES (?, ?, ?, ?)",
		Id_departures, Arrival, Id_route, Id_device)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetFlight(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `flight` "
	if filter != "" {
		query += " WHERE " + filter
	}

	query += ";"

	log.Print(query)

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag Flight
		selecte.Scan(&tag.Id, &tag.Id_departures, &tag.Arrival, &tag.Id_route, &tag.Id_device)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Id":             tag.Id,
			"Id_departurees": tag.Id_departures,
			"Arrival":        tag.Arrival,
			"Id_route":       tag.Id_route,
			"Id_device":      tag.Id_device,
		})
	}

	return return_val

}

func UpdateFlight(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `flight` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteFlight(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `flight` WHERE " + condition

	db.Query(query)

}

func GetFlightByCity(city string) []map[string]interface{} {
	type deviceType struct {
		types string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "SELECT device.type FROM `flight` JOIN `route` ON route.id = flight.id_route JOIN `device` ON device.id = flight.id_device WHERE route.arrival = \"" + city + "\";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag deviceType
		selecte.Scan(&tag.types)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"device ": tag.types,
		})
	}

	return return_val

}
