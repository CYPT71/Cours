package sql_request

import (
	"airfilgth/internal/utils"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Pilote struct {
	Id      int       `json:"id"`
	License time.Time `json:"licence"`
	StaffId int       `json:"staff_id"`
	Among   time.Time `json:"among"`
}

func AddPilote(license time.Time, among time.Time, staff_id int) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `Pilote`(`licence`, `among`, `staff_id`) VALUES  VALUES (?, ?, ?)",
		license, among, staff_id)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

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
		query += "*"
	}
	query += " FROM pilote "
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
		var tag Pilote
		selecte.Scan(&tag.Id, &tag.License, &tag.Among, &tag.StaffId)
		return_val = append(return_val, map[string]interface{}{
			"id":       tag.Id,
			"license":  tag.License.Format(time.UnixDate),
			"Among":    tag.Among.Format(time.UnixDate),
			"Staff id": tag.StaffId,
		})

	}
	return return_val

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