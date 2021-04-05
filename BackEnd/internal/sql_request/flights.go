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

	db.Exec("USE aircraft")

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

	db.Exec("USE aircraft")

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

	db.Exec("USE aircraft")

	query := "UPDATE `flight` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteFlight(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

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

	db.Exec("USE aircraft")

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

func OccupancyRate() []map[string]interface{} {

	type OccupancyRate struct {
		Arrival   string
		Occupancy int
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := `
		SELECT
			route.arrival,
			COUNT(route.arrival) AS "Occupancy rate"
		FROM
			departures
			JOIN tickets 
				ON tickets.departures_id = departures.id
			JOIN passenger 
				ON passenger.ticket_id = tickets.id
			JOIN flight 
				ON flight.id_departures = departures.id
			JOIN route 
				ON route.id = flight.id_route
		GROUP BY
			route.arrival
		ORDER BY
			COUNT(route.arrival)
		DESC
			;`

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag OccupancyRate
		selecte.Scan(&tag.Arrival, &tag.Occupancy)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Arrival":        tag.Arrival,
			"Occupancy Rate": tag.Occupancy,
		})
	}
	return result
}
