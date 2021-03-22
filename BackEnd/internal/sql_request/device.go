package sql_request

import (
	"airfilgth/internal/utils"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Device struct {
	Id       int    `json:"id"`
	Capacity int    `json:"capacity"`
	Type     string `json:"type"`
}

func AddDevices(capacity int, types string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `device` (`capacity`, `type`) VALUES (? , ?)", capacity, types)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetDevices(selector string, filter string) []map[string]interface{} {

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
	query += "FROM device "
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

func UpdateDevice(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	query := "UPDATE `device` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteDevice(condition string) {
	db, err := sql.Open("mysql", utils.Config.Mysql.Dns)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	query := "DELETE FROM `device` WHERE " + condition

	db.Query(query)

}
