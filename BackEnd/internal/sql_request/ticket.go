package sql_request

import (
	"airflight/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Tickets struct {
	Id            int
	Expire        string
	Price         int
	Departures_id int
}

func AddTickets(Expire string, Price int, Departures_id int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `tickets`(`expire`, `price`, `departures_id`) VALUES (?, ?, ?)",
		Expire, Price, Departures_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetTickets(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `tickets` "
	if filter != "" {

		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		return nil
	}

	var return_val []map[string]interface{}
	for selecte.Next() {
		var tag Tickets
		selecte.Scan(&tag.Id, &tag.Expire, &tag.Price, &tag.Departures_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Id":            tag.Id,
			"Expire":        tag.Expire,
			"Price":         tag.Price,
			"Departures_id": tag.Departures_id,
		})
	}

	return return_val
}

func UpdateTickets(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `tickets` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteTickets(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	query := "DELETE FROM `tickets` WHERE " + condition

	db.Query(query)

}

func TotalSales() int {

	type Sales struct {
		Price int
	}
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "SELECT SUM(price)  AS `total` FROM `tickets`;"

	selecte, err := db.Query(query)

	if err != nil {
		return -1
	}
	var result int
	for selecte.Next() {
		var tag Sales
		selecte.Scan(&tag.Price)

		result = tag.Price

	}

	return result
}

func SoldsPer(interval string) []map[string]interface{} {

	type Solds struct {
		Number string
		Date   string
	}

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var query string

	switch interval {
	default:
		query = `
		SELECT
			COUNT(tickets.id) as "Number",
			date_format(tickets.expire, '%Y-%m') as "month"
		FROM
			tickets
		GROUP by month;`
		break
	case "month":
		query = `
			SELECT
				COUNT(tickets.id) AS "Number",
				date_format(tickets.expire, '%Y-%m') as "month"
			FROM
				tickets
			GROUP by month;`
	case "day":
		query = `
			SELECT
    			COUNT(tickets.id) as "Number",
    			date_format(tickets.expire, '%Y-%m-%d') as "month"
			FROM
    			tickets
			GROUP by month;`
	case "week":
		query = `
			SELECT
				COUNT(tickets.id) AS "Number",
				STR_TO_DATE(CONCAT(YEARWEEK(tickets.expire),' Monday'), '%X%V %W') AS "week" 
			FROM
				tickets
			GROUP by STR_TO_DATE(CONCAT(YEARWEEK(tickets.expire),' Monday'), '%X%V %W');`

	}

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var result []map[string]interface{}

	for selecte.Next() {
		var tag Solds
		selecte.Scan(&tag.Number, &tag.Date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		result = append(result, map[string]interface{}{
			"Number": tag.Number,
			"Date":   tag.Date,
		})
	}
	return result
}
