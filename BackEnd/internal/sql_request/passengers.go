package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

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
