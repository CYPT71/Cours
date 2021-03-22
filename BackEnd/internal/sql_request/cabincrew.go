package sql_request

import (
	"airfilgth/internal/utils"
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CabinCrew struct {
	id       int
	fonction string
	among    time.Time
	staff_id int
}

func AddCabinCrew(fonction string, among time.Time, staff_id int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `cabincrew`(`fonction`, `among`, `staff_id`) VALUES  VALUES (?, ?, ?)",
		fonction, among, staff_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetCabinCrew(selector string, filter string) [][]string {

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

	var return_val [][]string
	var tag CabinCrew
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.fonction, &tag.among, &tag.staff_id)
		to_inject := []string{strconv.Itoa(tag.id), tag.fonction, tag.among.Format(time.UnixDate), strconv.Itoa(tag.staff_id)}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdateCabinCrew(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `cabincrew` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteCabinCrew(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `cabincrew` WHERE " + condition

	db.Query(query)

}
