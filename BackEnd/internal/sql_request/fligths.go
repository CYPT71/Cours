package sql_request

import (
	"airfilgth/internal/utils"
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Fligth struct {
	id            int
	id_departures int
	ariaval       time.Time
	id_route      int
	id_device     int
}

func AddFligth(id_departures int, ariaval time.Time, id_route int, id_device int) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Fligth`(`id_departures`, `arrival`, `id_route`, `id_device`) VALUES (?, ?, ?, ?)",
		id_departures, ariaval, id_route, id_device)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetFligth(selector string, filter string) [][]string {

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
	query += "FROM Fligth "
	if filter != "" {
		query += " WHERE " + filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val [][]string
	var tag Fligth
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.id_departures, &tag.id_device, &tag.id_route, &tag.ariaval)
		to_inject := []string{strconv.Itoa(tag.id), strconv.Itoa(tag.id_departures), strconv.Itoa(tag.id_device),
			strconv.Itoa(tag.id_route), tag.ariaval.Format(time.UnixDate)}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdateFligth(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `Fligth` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteFligth(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "DELETE FROM `Fligth` WHERE " + condition

	db.Query(query)

}
