package sql_request

import (
	"airflight/internal/utils"
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Passenger struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	First_name string `json:"first_name"`
	Address    string `json:"address"`
	Profession string `json:"profession"`
	Bank       string `json:"bank"`
	Ticket_id  int    `json:"ticket_id"`
}

func AddPassenger(Name string, First_name string, Address string, Profession string, Bank int, Ticket_id int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `passenger`(`name`, `first_name`, `address`, `profession`, `bank`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?)",
		Name, First_name, Address, Profession, Bank, Ticket_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetPassenger(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `passenger` "
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
		var tag Passenger
		selecte.Scan(&tag.Id, &tag.Name, &tag.First_name, &tag.Address, &tag.Profession, &tag.Bank, &tag.Ticket_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"id":          tag.Id,
			"Address":     tag.Address,
			"Bank":        tag.Bank,
			"First Name":  tag.First_name,
			"Name":        tag.Name,
			"Proffession": tag.Profession,
			"Ticket id":   tag.Ticket_id,
		})
	}

	return return_val

}

func UpdatePassenger(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "UPDATE `passenger` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeletePassenger(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "DELETE FROM `passenger` WHERE " + condition

	db.Query(query)

}

func ListPassengerperFlight() []map[string]interface{} {
	type QueryIdRoute struct {
		IdRoute int `json:"id_route"`
	}
	type Passengers struct {
		Name      string `json:"name"`
		FirstName string `json:"first_name"`
		Address   string `json:"address"`
		TicketId  string `json:"ticket_id"`
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "SELECT id_route from flight"
	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var idsRoute QueryIdRoute
		selecte.Scan(&idsRoute.IdRoute)
		query := "SELECT name, first_name, address, ticket_id FROM `flight` JOIN `tickets` ON tickets.departures_id = flight.id_departures JOIN `passenger` ON passenger.ticket_id = tickets.id WHERE id_route = " + strconv.Itoa(idsRoute.IdRoute)

		select_sub, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		var passenger []map[string]interface{}
		for select_sub.Next() {
			var getInfo Passengers
			select_sub.Scan(&getInfo.Name, &getInfo.FirstName, &getInfo.Address, &getInfo.TicketId)
			passenger = append(passenger, map[string]interface{}{
				"Name":       getInfo.Name,
				"First Name": getInfo.FirstName,
				"address":    getInfo.Address,
				"ticket id":  getInfo.TicketId,
			})
		}

		if err != nil {
			panic(err.Error())
		}

		result = append(result, map[string]interface{}{
			"Origin":    GetRoute("", "`id` ="+strconv.Itoa(idsRoute.IdRoute))[0]["Origin"],
			"Passenger": passenger,
		})

	}

	return result

}

func MostRegularProfession() []map[string]interface{} {
	type RegularProfession struct {
		Profession string
		Count      int
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := "SELECT profession, MAX(regular) \"passengers\" FROM (SELECT profession, COUNT(profession) AS \"regular\" FROM `passenger` GROUP BY profession) as tab1 GROUP BY profession ORDER BY profession DESC;"
	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var idsRoute RegularProfession
		selecte.Scan(&idsRoute.Profession, &idsRoute.Count)

		result = append(result, map[string]interface{}{
			"Profession": idsRoute.Profession,
			"Count":      idsRoute.Count,
		})

	}

	return result

}

func MostRegularPassenger() []map[string]interface{} {

	type MostRegular struct {
		Name          string
		FirstName     string
		NumberTickets int
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := `
		SELECT
			passenger.name,
			passenger.first_name,
			COUNT(tickets.id) as "c"
		FROM passenger
			JOIN tickets 
				ON tickets.id = passenger.ticket_id
		GROUP BY date_format(tickets.expire, '%Y-%m'), passenger.name, passenger.first_name
		HAVING c >= 2;`

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag MostRegular
		selecte.Scan(&tag.Name, &tag.FirstName, &tag.NumberTickets)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Name":                                 tag.Name,
			"First Name":                           tag.FirstName,
			"highest Number of tickets in a month": tag.NumberTickets,
		})
	}
	return result
}

func NumbOfPassengersByPeriodByPlane(start string, end string) []map[string]interface{} {

	type NumbPassengers struct {
		NumberOfPassengers int
		Date               string
		Type               string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := `
		SELECT DISTINCT
			departures.occupied AS "number of passengers transported by plane", departures.date, device.type
		FROM
			passenger
				JOIN tickets 
					ON passenger.ticket_id = tickets.id
				JOIN departures 
					ON tickets.departures_id = departures.id
				JOIN flight 
					ON departures.id = flight.id_departures
				JOIN device 
					ON device.id = flight.id_device
		WHERE
			departures.date BETWEEN `
	query += "\"" + start + "\" AND \"" + end + "\";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag NumbPassengers
		selecte.Scan(&tag.NumberOfPassengers, &tag.Date, &tag.Type)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Number of passengers": tag.NumberOfPassengers,
			"Date":                 tag.Date,
			"Type of the plane":    tag.Type,
		})
	}
	return result
}

func NumbOfPassengersByPeriod(start string, end string) []map[string]interface{} {

	type NumbPassengers struct {
		NumberOfPassengers int
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.Exec("USE aircraft")

	query := `
		SELECT
			SUM(departures.occupied) AS "number of passengers carried"
		FROM
			passenger
				JOIN tickets 
					ON passenger.ticket_id = tickets.id
				JOIN departures 
					ON tickets.departures_id = departures.id
				JOIN flight 
					ON departures.id = flight.id_departures
		WHERE
			departures.date BETWEEN `
	query += "\"" + start + "\" AND \"" + end + "\";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag NumbPassengers
		selecte.Scan(&tag.NumberOfPassengers)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Number of passengers": tag.NumberOfPassengers,
		})
	}
	return result
}

func AverageOccupancyRate(stage string) []map[string]interface{} {

	var query string
	var response []map[string]interface{}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.Exec("USE aircraft")

	switch stage {
	default:
		type ReturnStruct struct {
			DeviceId       int
			DeviceType     string
			DeviceCapacity int
			Average        string
		}
		query = `SELECT
			device.id,
			device.type,
			device.capacity,
			(
				SUM(
					departures.occupied / device.capacity
				) / COUNT(device.id)
			) AS "Occupancy rate by plane"
		FROM
			device
		LEFT JOIN flight ON flight.id_device = device.id
		JOIN departures ON departures.id = flight.id_departures
		GROUP BY
			device.id;`

		selecte, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		var result []map[string]interface{}

		for selecte.Next() {
			var tag ReturnStruct
			selecte.Scan(&tag.DeviceId, &tag.DeviceType, &tag.DeviceCapacity, &tag.Average)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			result = append(result, map[string]interface{}{
				"Device id":       tag.DeviceId,
				"Device Type":     tag.DeviceType,
				"Device Capacity": tag.DeviceCapacity,
				"Average":         tag.Average,
			})
		}

		response = result

		break
	case "flight":
		type ResponseStruct struct {
			FlightId   int
			DeviceType string
			FreePace   int
			Occupied   int
			Capacity   int
			Average    string
		}
		query = `
			SELECT
				flight.id,
				device.type,
				departures.free_places,
				departures.occupied,
				device.capacity,
				departures.occupied / device.capacity AS "Occupancy rate by flight"
			FROM
				flight
			JOIN device ON device.id = flight.id_device
			JOIN departures ON departures.id = flight.id_departures;
			`
		selecte, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		var result []map[string]interface{}

		for selecte.Next() {
			var tag ResponseStruct
			selecte.Scan(&tag.FlightId, &tag.DeviceType, &tag.FreePace, &tag.Occupied, &tag.Capacity, &tag.Average)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			result = append(result, map[string]interface{}{
				"Flight id":   tag.FlightId,
				"Device Type": tag.DeviceType,
				"Free Dents":  tag.FreePace,
				"ifrice":      tag.Occupied,
				"Capacity":    tag.Capacity,
				"Average":     tag.Average,
			})
		}

		response = result
		break
	case "destination":
		type ResponseStruct struct {
			Arrival string
			Average string
		}
		query =
			`
				SELECT
					arrival,
					(
						SUM(
							departures.occupied / device.capacity
						) / COUNT(route.arrival)
					) AS "Occupancy rate by destination"
				FROM
					route
				JOIN flight ON flight.id_route = route.id
				JOIN device ON device.id = flight.id_device
				JOIN departures ON departures.id = flight.id_departures
				GROUP BY
					arrival;`
		selecte, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		var result []map[string]interface{}

		for selecte.Next() {
			var tag ResponseStruct
			selecte.Scan(&tag.Arrival, &tag.Average)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			result = append(result, map[string]interface{}{
				"Arrival": tag.Arrival,
				"Average": tag.Average,
			})
		}

		response = result
		break
	case "plane":
		type ReturnStruct struct {
			DeviceId       int
			DeviceType     string
			DeviceCapacity int
			Average        string
		}
		query = `SELECT
			device.id,
			device.type,
			device.capacity,
			(
				SUM(
					departures.occupied / device.capacity
				) / COUNT(device.id)
			) AS "Occupancy rate by plane"
		FROM
			device
		LEFT JOIN flight ON flight.id_device = device.id
		JOIN departures ON departures.id = flight.id_departures
		GROUP BY
			device.id;`

		selecte, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}

		var result []map[string]interface{}

		for selecte.Next() {
			var tag ReturnStruct
			selecte.Scan(&tag.DeviceId, &tag.DeviceType, &tag.DeviceCapacity, &tag.Average)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			result = append(result, map[string]interface{}{
				"Device id":       tag.DeviceId,
				"Device Type":     tag.DeviceType,
				"Device Capacity": tag.DeviceCapacity,
				"Average":         tag.Average,
			})
		}

		response = result

		break
	}
	return response
}
