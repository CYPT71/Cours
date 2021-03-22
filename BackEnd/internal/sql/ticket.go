package sql

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tickets struct {
	id            int
	expire        time.Time
	price         int
	departures_id int
}

func AddTickets(expire time.Time, price int, departures_id int) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Tickets`(`expire`, `price`, `departures_id`) VALUES (?, ?, ?)",
		expire, price, departures_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetTickets(selector string, filter string) [][]string {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM Tickets "
	if filter != "" {

		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		return nil
	}

	var return_val [][]string
	var tag Tickets
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.expire, &tag.price)
		to_inject := []string{strconv.Itoa(tag.id), tag.expire.Format(time.UnixDate), strconv.Itoa(tag.price)}
		return_val = append(return_val, to_inject)
	}

	return return_val
}

func UpdateTickets(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "UPDATE `Tickets` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteTickets(condition string) {
	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "DELETE FROM `Tickets` WHERE " + condition

	db.Query(query)

}
