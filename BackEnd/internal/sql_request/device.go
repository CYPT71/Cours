package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	Id       int    `json:"id"`
	Capacity int    `json:"capacity"`
	Type     string `json:"type"`
}

func GetDevices(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `device` "
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
		var tag Device
		selecte.Scan(&tag.Id, &tag.Capacity, &tag.Type)
		return_val = append(return_val, map[string]interface{}{
			"id":       tag.Id,
			"capacity": tag.Capacity,
			"type":     tag.Type,
		})
	}

	return return_val

}

func AddDevices(capacity int, types string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `device` (`capacity`, `type`) VALUES (? , ?)", capacity, types)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func UpdateDevice(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "UPDATE `device` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteDevice(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "DELETE FROM `device` WHERE " + condition

	db.Query(query)

}

func DeviveHours() []map[string]interface{} {

	type devicehoursStruct struct {
		Types       string
		FlightHours string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "SELECT type, SEC_TO_TIME(SUM(TIME_TO_SEC(among))) AS \"flight hours\" FROM `pilote` JOIN `departures` ON pilote.id = departures.pilote JOIN `flight` ON departures.id = flight.id_departures JOIN `device` ON device.id = flight.id_device GROUP BY device.id;"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag devicehoursStruct
		selecte.Scan(&tag.Types, &tag.FlightHours)
		return_val = append(return_val, map[string]interface{}{
			"type":         tag.Types,
			"flight hours": tag.FlightHours,
		})
	}

	return return_val
}
