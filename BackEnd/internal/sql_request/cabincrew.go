package sql_request

import (
	"airflight/internal/utils"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CabinCrew struct {
	Id       int
	Among    time.Time
	Fonction string
	Staff_id int
}

func AddCabincrew(Among time.Time, Fonction string, Staff_id int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `cabincrew`(`among`, `fonction`, `staff_id`) VALUES  VALUES (?, ?, ?)",
		Among, Fonction, Staff_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetCabincrew(selector string, filter string) []map[string]interface{} {

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
	query += "FROM `cabincrew` "
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
		var tag CabinCrew
		selecte.Scan(&tag.Id, &tag.Among, &tag.Fonction, &tag.Staff_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		return_val = append(return_val, map[string]interface{}{
			"Id":       tag.Id,
			"Among":    tag.Among.Format(time.UnixDate),
			"Fonction": tag.Fonction,
			"Staff id": tag.Staff_id,
		})
	}

	return return_val

}

func UpdateCabincrew(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `cabincrew` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteCabincrew(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `cabincrew` WHERE " + condition

	db.Query(query)

}
