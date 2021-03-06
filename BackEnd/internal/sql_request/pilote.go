package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Pilote struct {
	Id       int    `json:"id"`
	License  string `json:"license"`
	Among    string `json:"among"`
	Staff_Id int    `json:"staff_id"`
}

func GetPilote(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `pilote` "
	if filter != "" {
		query += "WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag Pilote
		selecte.Scan(&tag.Id, &tag.License, &tag.Among, &tag.Staff_Id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Id":       tag.Id,
			"License":  tag.License,
			"Among":    tag.Among,
			"Staff id": tag.Staff_Id,
		})
	}

	return return_val

}

func AddPilote(license string, among string, staff_id int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if license == "" || among == "" {
		return
	}
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `pilote`(`license`, `among`, `staff_id`) VALUES (?, ?, ?)",
		license, among, staff_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func UpdatePilote(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `Pilote` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeletePilote(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `Pilote` WHERE " + condition

	db.Query(query)

}

func GetPiloteAmong() []map[string]interface{} {
	type PiloteDetails struct {
		Name      string
		FirstName string
		Among     string
	}
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "SELECT name, first_name, among FROM `pilote` JOIN `employees` ON pilote.staff_id = employees.id JOIN `departures` ON pilote.id = departures.pilote;"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag PiloteDetails
		selecte.Scan(&tag.Name, &tag.FirstName, &tag.Among)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Name":       tag.Name,
			"First Name": tag.FirstName,
			"Among":      tag.Among,
		})

	}

	return return_val
}

func GetPiloteDestination(name string, firstName string) []map[string]interface{} {

	type PiloteRoute struct {
		Name      string
		FirstName string
		Depart    string
		Arrival   string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "SELECT name, first_name, route.origin, route.arrival FROM `pilote` JOIN `employees` ON pilote.staff_id = employees.id JOIN `departures` ON pilote.id = departures.pilote JOIN `flight` ON departures.id = flight.id_departures JOIN `route` ON route.id = flight.id_route WHERE name = \"" + name + "\" AND first_name = \"" + firstName + "\";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag PiloteRoute
		selecte.Scan(&tag.Name, &tag.FirstName, &tag.Depart, &tag.Arrival)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Name":       tag.Name,
			"First Name": tag.FirstName,
			"Depart":     tag.Depart,
			"Arrival":    tag.Arrival,
		})
	}

	return result

}

func GetAverageFlight() []map[string]interface{} {

	type Average struct {
		Average string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := `
		SELECT 
			MAX(flight.id) / MAX(pilote.id) AS "Average" 
		FROM 
			pilote

				JOIN departures 
					ON departures.pilote = pilote.id

				JOIN flight 
					ON flight.id_departures = departures.id;`

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag Average
		selecte.Scan(&tag.Average)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Aerage flights per pilot": tag.Average,
		})
	}

	return result

}

func CityOfThePilot() []map[string]interface{} {

	type PilotByArrival struct {
		Name         string
		FirstName    string
		RouteArrival string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := `
		SELECT 
			name, 
			first_name, 
			route.arrival 
		FROM 
			pilote 
				JOIN employees 
					ON pilote.staff_id = employees.id 
				JOIN departures 
					ON pilote.id = departures.pilote 
				JOIN flight 
					ON departures.id = flight.id_departures 
				JOIN route 
					ON route.id = flight.id_route 
		WHERE employees.address LIKE CONCAT("%", route.arrival, "%")`

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag PilotByArrival
		selecte.Scan(&tag.Name, &tag.FirstName, &tag.RouteArrival)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Name":       tag.Name,
			"First Name": tag.FirstName,
			"Arrival":    tag.RouteArrival,
		})
	}
	return result
}
