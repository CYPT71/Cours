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
	Address    string `json:"adress"`
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

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `passenger`(`name`, `first_name`, `adress`, `profession`, `bank`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?)",
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

	query := "UPDATE `passenger` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeletePassenger(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `passenger` WHERE " + condition

	db.Query(query)

}

func ListPassengerperFlight() []map[string]interface{} {
	type queryIdRoute struct {
		IdRoute int    `json:"id_rotue"`
		Origin  string `json:"Origin"`
	}
	type passengers struct {
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

	query := "SELECT id_rotue from flight"
	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	/*
		SELECT name, first_name, address, ticket_id FROM flight

		JOIN tickets ON tickets.departures_id = flight.id_departures

		JOIN passenger ON passenger.ticket_id = tickets.id

		WHERE id_route = 35;

	*/
	var result []map[string]interface{}

	for selecte.Next() {
		var idsRoute queryIdRoute
		selecte.Scan(&idsRoute.IdRoute, &idsRoute.Origin)

		query := "SELECT name, first_name, address, ticket_id FROM flight JOIN tickets ON tickets.departures_id = flight.id_departures JOIN passenger ON passenger.ticket_id = tickets.id WHERE id_route = " + strconv.Itoa(idsRoute.IdRoute)
		sub_select, err := db.Query(query)

		var passenger []map[string]interface{}
		for sub_select.Next() {
			var getInfo passengers
			sub_select.Scan(&getInfo.Name, &getInfo.FirstName, &getInfo.Address, &getInfo.TicketId)
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
			"Origin":    idsRoute.Origin,
			"Passenger": passenger,
		})

	}

	return result

}
